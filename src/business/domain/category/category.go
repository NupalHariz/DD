package category

import (
	"context"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/log"
	"github.com/reyhanmichiels/go-pkg/sql"
)

type Interface interface {
	Create(ctx context.Context, param entity.CategoryInputParam) error
}

type category struct {
	db  sql.Interface
	log log.Interface
}

type InitParam struct {
	Db  sql.Interface
	Log log.Interface
}

func Init(param InitParam) Interface {
	return &category{
		db:  param.Db,
		log: param.Log,
	}
}

func (c *category) Create(ctx context.Context, param entity.CategoryInputParam) error {
	err := c.createSQL(ctx, param)
	if err != nil {
		return err
	}

	return nil
}
