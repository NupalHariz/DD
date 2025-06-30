package assignment

import (
	"context"
	"fmt"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
	"github.com/reyhanmichiels/go-pkg/v2/errors"
	"github.com/reyhanmichiels/go-pkg/v2/query"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

func (a *assignment) createSQL(ctx context.Context, param entity.AssignmentInputParam) error {
	a.log.Debug(ctx, fmt.Sprintf("create assignment with body: %v", param))

	tx, err := a.db.Leader().BeginTx(ctx, "txAssignment", sql.TxOptions{})
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("iNewAssignment", insertAssignment, param)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if rowCount < 1 {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no assignment created")
	}

	if err := tx.Commit(); err != nil {
		return errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return nil
}

func (a *assignment) updateSQL(ctx context.Context, updateParam entity.AssignmentUpdateParam, assignmentParam entity.AssignmentParam) error {
	a.log.Debug(ctx, fmt.Sprintf("update assignment with id %d and body: %v", assignmentParam.Id, updateParam))
	
	qb := query.NewSQLQueryBuilder(a.db, "param", "db", &assignmentParam.Option)

	queryExt, queryArgs, err := qb.BuildUpdate(&updateParam, &assignmentParam)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	tx, err := a.db.Leader().BeginTx(ctx, "txAssignment", sql.TxOptions{})
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.Exec("uAssignment", updateAssignment+queryExt, queryArgs...)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if rowCount < 1 {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no assignment updated")
	}

	if err := tx.Commit(); err != nil {
		return errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}


	return nil
}

func (a *assignment) getAllSQL(ctx context.Context, param entity.AssignmentParam) ([]entity.Assignment, error) {
	var assignments []entity.Assignment

	a.log.Debug(ctx, fmt.Sprintf("get all assignments with param: %v", param))

	qb := query.NewSQLQueryBuilder(a.db, "param", "db", &param.Option)

	queryExt, queryArgs, countExt, countArgs, err := qb.Build(&param)
	if err != nil {
		return assignments, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	rows, err := a.db.Query(ctx, "raAssignment", readAssignment+queryExt, queryArgs...)
	if err != nil {
		return assignments, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var assignment entity.Assignment

		err = rows.StructScan(&assignment)
		if err != nil {
			return assignments, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
		}

		assignments = append(assignments, assignment)
	}

	pg := entity.Pagination{
		CurrentPage: param.PaginationParam.Page,
		CurrentElements: int64(len(assignments)),
	}

	if !param.Option.DisableLimit && len(assignments) > 0 {
		err := a.db.Get(ctx, "pAssignments", countAssignments+countExt, &pg.TotalElements, countArgs...)
		if err != nil {
			return assignments, errors.NewWithCode(codes.CodeSQLRead, err.Error())
		}
	}

	pg.ProcessPagination(param.Limit)

	a.log.Debug(ctx, fmt.Sprintf("success to get assignment with param: %v", param))

	return assignments, nil
}
