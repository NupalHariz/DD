package rest

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/NupalHariz/DD/src/business/usecase"
	"github.com/NupalHariz/DD/src/utils/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/go-pkg/v2/appcontext"
	"github.com/reyhanmichiels/go-pkg/v2/auth"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
	"github.com/reyhanmichiels/go-pkg/v2/errors"
	"github.com/reyhanmichiels/go-pkg/v2/log"
	"github.com/reyhanmichiels/go-pkg/v2/parser"
	"github.com/reyhanmichiels/go-pkg/v2/rate_limiter"
)

var once = &sync.Once{}

type REST interface {
	Run()
}

type rest struct {
	http        *gin.Engine
	uc          *usecase.Usecases
	ginConfig   config.GinConfig
	log         log.Interface
	rateLimiter rate_limiter.Interface
	json        parser.JSONInterface
	auth        auth.Interface
}

type InitParam struct {
	Uc          *usecase.Usecases
	GinConfig   config.GinConfig
	Log         log.Interface
	RateLimiter rate_limiter.Interface
	Json        parser.JSONInterface
	Auth        auth.Interface
}

func Init(param InitParam) REST {
	var r rest
	once.Do(func() {
		// set up gin mode
		switch param.GinConfig.Mode {
		case gin.ReleaseMode:
			gin.SetMode(gin.ReleaseMode)
		case gin.DebugMode, gin.TestMode:
			gin.SetMode(gin.TestMode)
		default:
			gin.SetMode("")
		}

		// initialize struct
		httpServer := gin.New()

		r = rest{
			http:        httpServer,
			uc:          param.Uc,
			ginConfig:   param.GinConfig,
			log:         param.Log,
			rateLimiter: param.RateLimiter,
			json:        param.Json,
			auth:        param.Auth,
		}

		// Set CORS
		switch r.ginConfig.CORS.Mode {
		case "allowall":
			r.http.Use(cors.New(cors.Config{
				AllowAllOrigins: true,
				AllowHeaders:    []string{"*"},
				AllowMethods: []string{
					http.MethodHead,
					http.MethodGet,
					http.MethodPost,
					http.MethodPut,
					http.MethodPatch,
					http.MethodDelete,
				},
			}))
		default:
			r.http.Use(cors.New(cors.DefaultConfig()))
		}

		// Set Recovery
		r.http.Use(r.CustomRecovery)

		// Set Timeout
		r.http.Use(r.SetTimeout)

		// TODO: set audit

		r.Register()
	})

	return &r
}

func (r *rest) Register() {
	// utility route
	r.http.GET("/ping", r.Ping)
	r.registerSwaggerRoutes()
	r.registerDummyRoutes()

	// Set Common Middlewares
	commonPublicMiddlewares := gin.HandlersChain{
		r.rateLimiter.Limiter(), r.addFieldsToContext, r.BodyLogger,
	}
	commonPrivateMiddlewares := gin.HandlersChain{
		r.rateLimiter.Limiter(), r.addFieldsToContext, r.BodyLogger, r.VerifyUser,
	}

	// auth api
	authV1 := r.http.Group("/auth/v1", commonPublicMiddlewares...)
	authV1.POST("/register", r.RegisterNewUser)
	authV1.POST("/login", r.SignInWithPassword)
	authV1.POST("/token/refresh", r.RefreshToken)

	// public api
	r.http.Group("/public/v1/", commonPublicMiddlewares...)

	// private api
	v1 := r.http.Group("/v1/", commonPrivateMiddlewares...)

	// category api
	v1.POST("/categories/", r.CreateCategory)
	v1.GET("/categories/", r.GetAllCategory)

	// budget api
	// v1.POST("/budgets/", r.CreateBudget)
	v1.PUT("/budgets/:id", r.UpdateBudget)
	v1.GET("/budgets/", r.GetAllBudget)

	// money api
	v1.POST("/moneys/", r.AddTransaction)
	v1.PUT("/moneys/:id", r.UpdateTransaction)
	v1.GET("/moneys", r.GetTransaction)

	// daily assignment api
	v1.POST("/daily-assignments/", r.CreateDailyAssignment)
	v1.PUT("/daily-assignments/:id", r.UpdateDailyAssignment)
	v1.GET("/daily-assignments", r.GetAllDailyAssignment)

	// assignment category
	v1.POST("/assignment-categories/", r.CreateAssignmentCategory)
	v1.GET("/assignment-categories/", r.GetAllAssignmentCategory)

	// assignment
	v1.POST("/assignments/", r.CreateAssignment)
	v1.PUT("/assignments/:id", r.UpdateAssignment)
	v1.GET("/assignments/", r.GetAllAssignment)
}

func (r *rest) Run() {
	// Create context that listens for the interrupt signal from the OS.
	c := appcontext.SetServiceVersion(context.Background(), r.ginConfig.Meta.Version)

	ctx, stop := signal.NotifyContext(c, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// configure server
	port := ":8080"
	if r.ginConfig.Port != "" {
		port = fmt.Sprintf(":%s", r.ginConfig.Port)
	}

	srv := &http.Server{
		Addr:              port,
		Handler:           r.http,
		ReadHeaderTimeout: 2 * time.Second,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			r.log.Error(ctx, fmt.Sprintf("Serving HTTP error: %s", err.Error()))
		}
	}()
	r.log.Info(ctx, fmt.Sprintf("Listening and Serving HTTP on %s", srv.Addr))

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	r.log.Info(ctx, "Shutting down server...")

	// The context is used to inform the server it has timeout duration to finish
	// the request it is currently handling
	quitContext, cancel := context.WithTimeout(c, r.ginConfig.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(quitContext); err != nil {
		r.log.Fatal(quitContext, fmt.Sprintf("Server Shutdown: %s", err.Error()))
	}

	r.log.Info(quitContext, "Server Shut Down.")
}

func (r *rest) Ping(ctx *gin.Context) {
	r.httpRespSuccess(ctx, codes.CodeSuccess, "ping success", nil)
}

func (r *rest) registerDummyRoutes() {
	if r.ginConfig.Dummy.Enabled {
		// load login page to gin
		r.http.LoadHTMLFiles(
			"./docs/templates/login.html",
		)

		dummyGroup := r.http.Group(r.ginConfig.Dummy.Path)
		dummyGroup.GET("/login", r.DummyLogin)
	}
}
