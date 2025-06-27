package usecase

import (
	"github.com/NupalHariz/DD/src/business/usecase/budget"
	"github.com/NupalHariz/DD/src/business/usecase/category"
	"github.com/NupalHariz/DD/src/business/usecase/money"

	"github.com/NupalHariz/DD/src/business/domain"
	"github.com/NupalHariz/DD/src/business/usecase/user"
	"github.com/reyhanmichiels/go-pkg/v2/auth"
	"github.com/reyhanmichiels/go-pkg/v2/hash"
	"github.com/reyhanmichiels/go-pkg/v2/log"
	"github.com/reyhanmichiels/go-pkg/v2/parser"
)

type Usecases struct {
	User     user.Interface
	Category category.Interface
	Budget   budget.Interface
	Money    money.Interface
}

type InitParam struct {
	Dom  *domain.Domains
	Json parser.JSONInterface
	Log  log.Interface
	Hash hash.Interface
	Auth auth.Interface
}

func Init(param InitParam) *Usecases {
	return &Usecases{
		User:     user.Init(user.InitParam{UserDomain: param.Dom.User, Auth: param.Auth, Hash: param.Hash}),
		Category: category.Init(category.InitParam{CategoryDom: param.Dom.Category, BudgetDom: param.Dom.Budget, Auth: param.Auth}),
		Budget:   budget.Init(budget.InitParam{Auth: param.Auth, BudgetDom: param.Dom.Budget, HistoryBudgetDom: param.Dom.HistoryBudget}),
		Money:    money.Init(money.InitParam{Auth: param.Auth, MoneyDom: param.Dom.Money, BudgetDom: param.Dom.Budget}),
	}
}
