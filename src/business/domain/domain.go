package domain

import (
	"github.com/NupalHariz/DD/src/business/domain/assignment"
	assignmentcategory "github.com/NupalHariz/DD/src/business/domain/assignment_category"
	"github.com/NupalHariz/DD/src/business/domain/budget"
	"github.com/NupalHariz/DD/src/business/domain/category"
	dailyassignment "github.com/NupalHariz/DD/src/business/domain/daily_assignment"
	historybudget "github.com/NupalHariz/DD/src/business/domain/history_budget"
	"github.com/NupalHariz/DD/src/business/domain/money"

	"github.com/NupalHariz/DD/src/business/domain/user"
	"github.com/reyhanmichiels/go-pkg/v2/log"
	"github.com/reyhanmichiels/go-pkg/v2/parser"
	"github.com/reyhanmichiels/go-pkg/v2/redis"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

type Domains struct {
	User               user.Interface
	Category           category.Interface
	Budget             budget.Interface
	Money              money.Interface
	HistoryBudget      historybudget.Interface
	DailyAssignment    dailyassignment.Interface
	AssignmentCategory assignmentcategory.Interface
	Assignment         assignment.Interface
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
		User:               user.Init(user.InitParam{Db: param.Db, Log: param.Log, Redis: param.Redis, Json: param.Json}),
		Category:           category.Init(category.InitParam{Db: param.Db, Log: param.Log, Json: param.Json, Redis: param.Redis}),
		Budget:             budget.Init(budget.InitParam{Db: param.Db, Log: param.Log, Json: param.Json, Redis: param.Redis}),
		Money:              money.Init(money.InitParam{Db: param.Db, Log: param.Log, Json: param.Json, Redis: param.Redis}),
		HistoryBudget:      historybudget.Init(historybudget.InitParam{Db: param.Db, Log: param.Log}),
		DailyAssignment:    dailyassignment.Init(dailyassignment.InitParam{Db: param.Db, Log: param.Log, Json: param.Json, Redis: param.Redis}),
		AssignmentCategory: assignmentcategory.Init(assignmentcategory.InitParam{Db: param.Db, Log: param.Log, Json: param.Json, Redis: param.Redis}),
		Assignment:         assignment.Init(assignment.InitParam{Db: param.Db, Log: param.Log, Json: param.Json, Redis: param.Redis}),
	}
}
