package dailyassignment

import (
	"context"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/log"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

type Interface interface {
	Create(ctx context.Context, param entity.DailyAssignmentInputParam) error
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
