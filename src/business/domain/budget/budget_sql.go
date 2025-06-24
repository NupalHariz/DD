package budget

import (
	"context"
	"fmt"
	"strings"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
	"github.com/reyhanmichiels/go-pkg/v2/errors"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

func (b *budget) CreateSQL(ctx context.Context, param entity.BudgetInputParam) error {
	b.log.Info(ctx, fmt.Sprintf("create budget with body: %v", param))

	tx, err := b.db.Leader().BeginTx(ctx, "txBudget", sql.TxOptions{})
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("iNewCategory", insertBudget, param)
	if err != nil {
		if strings.Contains(err.Error(), entity.DuplicateEntryErrMessage) {
			return errors.NewWithCode(codes.CodeSQLUniqueConstraint, err.Error())
		}

		return errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if rowCount < 1 {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no budget created")
	}

	if err := tx.Commit(); err != nil {
		return errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return nil
}

func (b *budget) updateExpenseSQL(ctx context.Context, updateParam entity.BudgetUpdateParam) error {
	b.log.Info(ctx, fmt.Sprintf(
		"adding %d into current expense with user_id = %d and category_id = %d",
		updateParam.CurrentExpense,
		updateParam.UserId,
		updateParam.CategoryId),
	)

	tx, err := b.db.Leader().BeginTx(ctx, "txBudget", sql.TxOptions{})
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("uBudgetExpense", updateCurrentExpense, updateParam)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	rowCount, err := res.RowsAffected()

	if err != nil {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if rowCount < 1 {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no budget updated")
	}

	if err := tx.Commit(); err != nil {
		return errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}
	return nil
}
