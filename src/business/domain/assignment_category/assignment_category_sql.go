package assignmentcategory

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

func (a *assignmentCategory) createSQL(ctx context.Context, param entity.AssignmentCategoryInputParam) error {
	a.log.Debug(ctx, fmt.Sprintf("create assignment category with body: %v", param))

	tx, err := a.db.Leader().BeginTx(ctx, "txAssignmentCategory", sql.TxOptions{})
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("iNewAssignmentCategory", insertAssignmentCategory, param)
	if err != nil {
		if strings.Contains(err.Error(), entity.DuplicateEntryErrMessage) {
			return errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
		}

		return errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if rowCount < 1 {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no assignment category created")
	}

	if err := tx.Commit();  err != nil {
		return errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}
	
	return nil
}

func (a *assignmentCategory) getAllSQL(ctx context.Context, param entity.AssignmentCategoryParam) ([]entity.AssignmentCategory, error) {
	var assignmentCategories []entity.AssignmentCategory

	a.log.Debug(ctx, fmt.Sprintf("get all assignment category with param: %v", param))

	qb := query.NewSQLQueryBuilder(a.db, "param", "db", &param.Option)

	queryExt, queryArgs, _, _, err := qb.Build(&param)
	if err != nil {
		return assignmentCategories, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	rows, err := a.db.Query(ctx, "reAssignmentCategory", readAssignmentCategory+queryExt, queryArgs...)
	if err != nil {
		return assignmentCategories, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}
	defer rows.Close()

	for rows.Next(){
		var assignmentCategory entity.AssignmentCategory

		err = rows.StructScan(&assignmentCategory)
		if err != nil {
			return assignmentCategories, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
		}

		assignmentCategories = append(assignmentCategories, assignmentCategory)
	}

	return assignmentCategories, nil
}