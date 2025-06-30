package dailyassignment

import (
	"context"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/log"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

type Interface interface {
	Create(ctx context.Context, param entity.DailyAssignmentInputParam) error
	Update(ctx context.Context, updateParam entity.DailyAssignmentUpdateParam, dailyAssignmentParam entity.DailyAssignmentParam) error
	GetAll(ctx context.Context, param entity.DailyAssignmentParam) ([]entity.DailyAssignment, error)
}

type dailyAssignment struct {
	db  sql.Interface
	log log.Interface
}

type InitParam struct {
	Db  sql.Interface
	Log log.Interface
}

func Init(param InitParam) Interface {
	return &dailyAssignment{
		db:  param.Db,
		log: param.Log,
	}
}

func (d *dailyAssignment) Create(ctx context.Context, param entity.DailyAssignmentInputParam) error {
	err := d.createSQL(ctx, param)
	if err != nil {
		return err
	}

	return nil
}

func (d *dailyAssignment) Update(
	ctx context.Context,
	updateParam entity.DailyAssignmentUpdateParam,
	dailyAssignmentParam entity.DailyAssignmentParam,
) error {
	err := d.updateSQL(ctx, updateParam, dailyAssignmentParam)
	if err != nil {
		return err
	}

	return nil
}

func (d *dailyAssignment) GetAll(ctx context.Context, param entity.DailyAssignmentParam) ([]entity.DailyAssignment, error) {
	data, err := d.getAllSQL(ctx, param)
	if err != nil {
		return []entity.DailyAssignment{}, err
	}

	return data, nil
}
