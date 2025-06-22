package domain

import (
	"github.com/NupalHariz/DD/src/business/domain/budget"
	"github.com/NupalHariz/DD/src/business/domain/category"

	"github.com/NupalHariz/DD/src/business/domain/user"
	"github.com/reyhanmichiels/go-pkg/log"
	"github.com/reyhanmichiels/go-pkg/parser"
	"github.com/reyhanmichiels/go-pkg/redis"
	"github.com/reyhanmichiels/go-pkg/sql"
)

type Domains struct {
	User     user.Interface
	Category category.Interface
	Budget   budget.Interface
}

type InitParam struct {
	Log   log.Interface
	Db    sql.Interface
	Redis redis.Interface
	Json  parser.JSONInterface
	// TODO: add audit
}

func Init(param InitParam) *Domains {
	return &Domains{
		User:     user.Init(user.InitParam{Db: param.Db, Log: param.Log, Redis: param.Redis, Json: param.Json}),
		Category: category.Init(category.InitParam{Db: param.Db, Log: param.Log}),
		Budget:   budget.Init(budget.InitParam{Db: param.Db, Log: param.Log}),
	}
}
