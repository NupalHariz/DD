package historybudget

import (
	"context"
	"fmt"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
	"github.com/reyhanmichiels/go-pkg/v2/errors"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

func (h *historyBudget) createBatchSQL(ctx context.Context, params []entity.HistoryBudget) error {
	h.log.Debug(ctx, fmt.Sprintf("create history_budgets with body: %v", params))

	tx, err := h.db.Leader().BeginTx(ctx, "txHistoryBudget", sql.TxOptions{})
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("iHistoryBudget", insertHistoryBudget, params)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if rowCount < 1 {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no history budget created")
	}

	if err := tx.Commit(); err != nil {
		return errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	h.log.Debug(ctx, fmt.Sprintf("success to create history budgets with body: %v", params))
	
	return nil
}
