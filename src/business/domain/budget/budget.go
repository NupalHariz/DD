package budget

import (
	"context"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/log"
	"github.com/reyhanmichiels/go-pkg/sql"
)

type Interface interface {
	Create(ctx context.Context, param entity.BudgetInputParam) error
}

type budget struct {
	db  sql.Interface
	log log.Interface
}

type InitParam struct {
	Db  sql.Interface
	Log log.Interface
}

func Init(param InitParam) Interface {
	return &budget{
		db:  param.Db,
		log: param.Log,
	}
}

func (b *budget) Create(ctx context.Context, param entity.BudgetInputParam) error {
	err := b.CreateSQL(ctx, param)
	if err != nil {
		return err
	}
	return nil
}
