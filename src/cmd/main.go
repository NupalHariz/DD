package main

import (
	"context"
	"errors"

	"github.com/NupalHariz/DD/src/business/domain"
	"github.com/NupalHariz/DD/src/business/service/mail"
	"github.com/NupalHariz/DD/src/business/usecase"
	"github.com/NupalHariz/DD/src/handler/rest"
	"github.com/NupalHariz/DD/src/handler/scheduler"
	"github.com/NupalHariz/DD/src/utils/config"
	"github.com/reyhanmichiels/go-pkg/v2/auth"
	"github.com/reyhanmichiels/go-pkg/v2/configreader"
	"github.com/reyhanmichiels/go-pkg/v2/files"
	"github.com/reyhanmichiels/go-pkg/v2/hash"
	"github.com/reyhanmichiels/go-pkg/v2/log"
	"github.com/reyhanmichiels/go-pkg/v2/parser"
	"github.com/reyhanmichiels/go-pkg/v2/rate_limiter"
	"github.com/reyhanmichiels/go-pkg/v2/redis"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

// @contact.name   Naufal Haris Rusyard
// @contact.email  naufal.michiels@gmail.com

// @securitydefinitions.apikey BearerAuth
// @in header
// @name Authorization

const (
	configfile   string = "./etc/cfg/conf.json"
	templatefile string = "./etc/tpl/conf.template.json"
	appnamespace string = ""
)

func main() {
	defaultLogger := log.DefaultLogger()

	// panic recovery
	defer func() {
		if err := recover(); err != nil {
			defaultLogger.Panic(err)
		}
	}()

	// TODO: need a way to build config file automatically, for now build the file manually
	if !files.IsExist(configfile) {
		defaultLogger.Fatal(context.Background(), errors.New("config file doesn't exist"))
	}

	// read config from config file
	cfg := config.Init()
	configReader := configreader.Init(configreader.Options{
		ConfigFile: configfile,
	})
	configReader.ReadConfig(&cfg)

	// init logger
	log := log.Init(cfg.Log)

	// init cache
	cache := redis.Init(cfg.Redis, log)

	// init db
	db := sql.Init(cfg.SQL, log)

	// init rate limiter
	rateLimiter := rate_limiter.Init(cfg.RateLimiter, log)

	// init parser
	parser := parser.InitParser(log, cfg.Parser)

	// init domain
	dom := domain.Init(domain.InitParam{Log: log, Db: db, Redis: cache, Json: parser.JSONParser()})

	// hash
	hash := hash.Init()

	// auth
	auth := auth.Init(cfg.Auth, log)

	mail := mail.Init(mail.InitParam{Log: log, Cfg: cfg.Mail})

	// init usecase
	uc := usecase.Init(usecase.InitParam{Dom: dom, Log: log, Json: parser.JSONParser(), Hash: hash, Auth: auth, Mail: mail})

	sch := scheduler.Init(
		scheduler.InitParam{
			MetaConf: cfg.Meta,
			Log: log,
			Uc: uc,
		},
	)

	// init http server
	r := rest.Init(rest.InitParam{Uc: uc, GinConfig: cfg.Gin, Log: log, RateLimiter: rateLimiter, Json: parser.JSONParser(), Auth: auth})

	sch.Run()

	// run http server
	r.Run()
}
