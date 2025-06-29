package category

import (
	"context"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/log"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

type Interface interface {
	Create(ctx context.Context, param entity.CategoryInputParam) (entity.Category, error)
	GetAll(ctx context.Context, param entity.CategoryParam) ([]entity.Category, error)
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

func (c *category) Create(ctx context.Context, param entity.CategoryInputParam) (entity.Category, error) {
	category, err := c.createSQL(ctx, param)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (c *category) GetAll(ctx context.Context, param entity.CategoryParam) ([]entity.Category, error) {
	categories, err := c.getAllSQL(ctx, param)
	if err != nil {
		return categories, err
	}

	return categories, nil
}
