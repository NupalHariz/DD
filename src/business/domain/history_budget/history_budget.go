package historybudget

import (
	"context"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/log"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

type Interface interface {
	CreateBatch(ctx context.Context, params []entity.HistoryBudget) error
}

type historyBudget struct {
	db  sql.Interface
	log log.Interface
}

type InitParam struct {
	Db  sql.Interface
	Log log.Interface
}

func Init(param InitParam) Interface {
	return &historyBudget{
		db:  param.Db,
		log: param.Log,
	}
}

func (h *historyBudget) CreateBatch(ctx context.Context, params []entity.HistoryBudget) error {
	err := h.createBatchSQL(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
