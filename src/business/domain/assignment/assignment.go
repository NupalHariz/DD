package assignment

import (
	"context"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/log"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

type Interface interface {
	Create(ctx context.Context, param entity.AssignmentInputParam) error
	Update(ctx context.Context, updateParam entity.AssignmentUpdateParam, assignmentParam entity.AssignmentParam) error
}

type assignment struct {
	db  sql.Interface
	log log.Interface
}

type InitParam struct {
	Db  sql.Interface
	Log log.Interface
}

func Init(param InitParam) Interface {
	return &assignment{
		db:  param.Db,
		log: param.Log,
	}
}

func (a *assignment) Create(ctx context.Context, param entity.AssignmentInputParam) error {
	err := a.createSQL(ctx, param)
	if err != nil {
		return err
	}

	return nil
}

func (a *assignment) Update(ctx context.Context, updateParam entity.AssignmentUpdateParam, assignmentParam entity.AssignmentParam) error {
	err := a.updateSQL(ctx, updateParam, assignmentParam)
	if err != nil {
		return err
	}

	return nil
}
