package dailyassignment

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

func (d *dailyAssignment) createSQL(ctx context.Context, param entity.DailyAssignmentInputParam) error {
	d.log.Debug(ctx, fmt.Sprintf("creating daily assignment with body: %v", param))

	tx, err := d.db.Leader().BeginTx(ctx, "txDailyAssignment", sql.TxOptions{})
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("iNewDailyAssignment", insertDailyAssignment, param)
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
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no daily assignment created")
	}

	if err := tx.Commit(); err != nil {
		return errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return nil
}

func (d *dailyAssignment) updateSQL(
	ctx context.Context,
	updateParam entity.DailyAssignmentUpdateParam,
	dailyAssignmentParam entity.DailyAssignmentParam,
) error {
	d.log.Debug(ctx, fmt.Sprintf("update daily assignment with id %v and body %v", dailyAssignmentParam.Id, updateParam))

	qb := query.NewSQLQueryBuilder(d.db, "param", "db", &dailyAssignmentParam.Option)

	queryUpdate, queryArgs, err := qb.BuildUpdate(&updateParam, &dailyAssignmentParam)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	tx, err := d.db.Leader().BeginTx(ctx, "txDailyAssignment", sql.TxOptions{})
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.Exec("uDailyAssignment", updateDailyAssignment+queryUpdate, queryArgs...)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if rowCount < 1 {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no daily assignment updated")
	}

	if err := tx.Commit(); err != nil {
		return errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return nil
}
