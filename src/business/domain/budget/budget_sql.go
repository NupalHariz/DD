package budget

import (
	"context"
	"fmt"
	"strings"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
	"github.com/reyhanmichiels/go-pkg/v2/errors"
	"github.com/reyhanmichiels/go-pkg/v2/query"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

func (b *budget) CreateSQL(ctx context.Context, param entity.BudgetInputParam) error {
	b.log.Debug(ctx, fmt.Sprintf("create budget with body: %v", param))

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

	b.log.Debug(ctx, fmt.Sprintf("success to create budget with body: %v", param))

	return nil
}

func (b *budget) updateExpenseSQL(ctx context.Context, updateParam entity.BudgetUpdateParam) error {
	b.log.Debug(ctx, fmt.Sprintf(
		"adding %v into current expense with user_i  %d and category_id %d",
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

	b.log.Debug(ctx, fmt.Sprintf("success to update expense with body: %v", updateParam))

	return nil
}

func (b *budget) updateSQL(ctx context.Context, updateParam entity.BudgetUpdateParam, budgetParam entity.BudgetParam) error {
	b.log.Debug(ctx, fmt.Sprintf("update budget with body %v and param %v", updateParam, budgetParam))

	qb := query.NewSQLQueryBuilder(b.db, "param", "db", &budgetParam.Option)

	queryUpdate, args, err := qb.BuildUpdate(&updateParam, &budgetParam)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	tx, err := b.db.Leader().BeginTx(ctx, "txBudget", sql.TxOptions{})
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.Exec("uBudget", updateBudget+queryUpdate, args...)
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

	b.log.Debug(ctx, fmt.Sprintf("success to update budget with body: %v", updateParam))

	return nil
}

func (b budget) getAllSQL(ctx context.Context, budgetParam entity.BudgetParam) ([]entity.Budget, error) {
	var budgets []entity.Budget

	b.log.Debug(ctx, fmt.Sprintf("get all budget with param: %v", budgetParam))

	qb := query.NewSQLQueryBuilder(b.db, "param", "db", &budgetParam.Option)

	queryExt, queryArgs, _, _, err := qb.Build(&budgetParam)
	if err != nil {
		return budgets, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	rows, err := b.db.Query(ctx, "rBudgetAll", readBudget+queryExt, queryArgs...)
	if err != nil {
		return budgets, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		budget := entity.Budget{}

		err := rows.StructScan(&budget)
		if err != nil {
			b.log.Error(ctx, codes.CodeSQLRowScan)
			continue
		}

		budgets = append(budgets, budget)
	}

	b.log.Debug(ctx, fmt.Sprintf("success to get budgets with param: %v", budgetParam))

	return budgets, err
}
