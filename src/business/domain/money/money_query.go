package money

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

func (m *money) createSql(ctx context.Context, param entity.MoneyInputParam) error {
	m.log.Debug(ctx, fmt.Sprintf("create money with body: %v", param))

	tx, err := m.db.Leader().BeginTx(ctx, "txMoney", sql.TxOptions{})
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("iNewMoney", insertMoney, param)
	if err != nil {
		if strings.Contains(err.Error(), entity.DuplicateEntryErrMessage) {
			return errors.NewWithCode(codes.CodeSQLUniqueConstraint, err.Error())
		}

		return errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	row, err := res.RowsAffected()
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if row < 1 {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no money created")
	}

	if err := tx.Commit(); err != nil {
		return errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}
	
	m.log.Debug(ctx, fmt.Sprintf("success to create money with body: %v", param))

	return nil
}

func (m *money) getSQL(ctx context.Context, param entity.MoneyParam) (entity.Money, error) {
	money := entity.Money{}

	m.log.Debug(ctx, fmt.Sprintf("read money with param: %v", param))

	qb := query.NewSQLQueryBuilder(m.db, "param", "db", &param.Option)

	queryExt, queryArgs, _, _, err := qb.Build(&param)
	if err != nil {
		return money, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	rows, err := m.db.QueryRow(ctx, "rMoney", readMoney+queryExt, queryArgs...)
	if err != nil {
		if errors.Is(err, sql.ErrNotFound) {
			return money, errors.NewWithCode(codes.CodeSQLRecordDoesNotExist, err.Error())
		}

		return money, errors.NewWithCode(codes.CodeSQL, err.Error())
	}

	if err := rows.StructScan(&money); err != nil {
		if errors.Is(err, sql.ErrNotFound) {
			return money, errors.NewWithCode(codes.CodeSQLRecordDoesNotExist, err.Error())
		}

		return money, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	}

	m.log.Debug(ctx, fmt.Sprintf("success to get money with body: %v", money))

	return money, nil
}

func (m *money) updateSQL(ctx context.Context, updateParam entity.MoneyUpdateParam, moneyParam entity.MoneyParam) error {
	m.log.Debug(ctx, fmt.Sprintf("update money with body %v and param %v", updateParam, moneyParam))

	qb := query.NewSQLQueryBuilder(m.db, "param", "db", &moneyParam.Option)

	queryUpdate, args, err := qb.BuildUpdate(&updateParam, &moneyParam)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	tx, err := m.db.Leader().BeginTx(ctx, "txMoney", sql.TxOptions{})
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.Exec("uMoney", updateMoney+queryUpdate, args...)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if rowCount < 1 {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no money updated")
	}

	if err := tx.Commit(); err != nil {
		return errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	m.log.Debug(ctx, fmt.Sprintf("success to upgrade money with body: %v", updateParam))

	return nil
}
