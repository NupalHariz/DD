package usecase

import (
	"github.com/NupalHariz/DD/src/business/usecase/category"

	"github.com/NupalHariz/DD/src/business/domain"
	"github.com/NupalHariz/DD/src/business/usecase/user"
	"github.com/reyhanmichiels/go-pkg/auth"
	"github.com/reyhanmichiels/go-pkg/hash"
	"github.com/reyhanmichiels/go-pkg/log"
	"github.com/reyhanmichiels/go-pkg/parser"
)

type Usecases struct {
	User     user.Interface
	Category category.Interface
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
		Category: category.Init(category.InitParam{CategoryDom: param.Dom.Category, Auth: param.Auth}),
	}
}
