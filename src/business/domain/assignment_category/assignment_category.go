package assignmentcategory

import (
	"context"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/log"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

type Interface interface {
	Create(ctx context.Context, param entity.AssignmentCategoryInputParam) error
	GetAll(ctx context.Context, param entity.AssignmentCategoryParam) ([]entity.AssignmentCategory, error)
}

type assignmentCategory struct {
	db  sql.Interface
	log log.Interface
}

type InitParam struct {
	Db  sql.Interface
	Log log.Interface
}

func Init(param InitParam) Interface {
	return &assignmentCategory{
		db:  param.Db,
		log: param.Log,
	}
}

func (a *assignmentCategory) Create(ctx context.Context, param entity.AssignmentCategoryInputParam) error {
	err := a.createSQL(ctx, param)
	if err != nil {
		return err
	}

	return nil
}

func (a *assignmentCategory) GetAll(ctx context.Context, param entity.AssignmentCategoryParam) ([]entity.AssignmentCategory, error) {
	categories, err := a.getAllSQL(ctx, param)
	if err != nil {
		return []entity.AssignmentCategory{}, err
	}

	return categories, nil
}
