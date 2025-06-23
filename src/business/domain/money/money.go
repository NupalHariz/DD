package money

import (
	"context"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/log"
	"github.com/reyhanmichiels/go-pkg/sql"
)

type Interface interface {
	Create(ctx context.Context, param entity.MoneyInputParam) error
}

type money struct {
	db  sql.Interface
	log log.Interface
}

type InitParam struct {
	Db  sql.Interface
	Log log.Interface
}

func Init(param InitParam) Interface {
	return &money{
		db:  param.Db,
		log: param.Log,
	}
}

func (m *money) Create(ctx context.Context, param entity.MoneyInputParam) error {
	err := m.createSql(ctx, param)
	if err != nil {
		return err
	}

	return nil
}
