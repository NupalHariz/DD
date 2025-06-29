package category

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

func (c *category) createSQL(ctx context.Context, param entity.CategoryInputParam) (entity.Category, error) {
	var category entity.Category

	c.log.Debug(ctx, fmt.Sprintf("create category with body: %v", param))

	tx, err := c.db.Leader().BeginTx(ctx, "txCategory", sql.TxOptions{})
	if err != nil {
		return category, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("iNewCategory", insertCategory, param)
	if err != nil {
		if strings.Contains(err.Error(), entity.DuplicateEntryErrMessage) {
			return category, errors.NewWithCode(codes.CodeSQLUniqueConstraint, err.Error())
		}

		return category, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return category, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if rowCount < 1 {
		return category, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no category created")
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return category, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}

	if err := tx.Commit(); err != nil {
		return category, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	c.log.Debug(ctx, fmt.Sprintf("success create user with body: %v", param))

	category = entity.Category{
		Id:     lastId,
		UserId: param.UserId,
		Name:   param.Name,
	}

	return category, nil
}

func (c *category) getAllSQL(ctx context.Context, param entity.CategoryParam) ([]entity.Category, error) {
	var categories []entity.Category

	c.log.Debug(ctx, fmt.Sprintf("read all category with param: %v", param))

	qb := query.NewSQLQueryBuilder(c.db, "param", "db", &param.Option)

	queryExt, queryArgs, _, _, err := qb.Build(&param)
	if err != nil {
		return categories, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	rows, err := c.db.Query(ctx, "raCategory", readCategories+queryExt, queryArgs...)
	if err != nil {
		return categories, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	for rows.Next(){
		var category entity.Category
		err := rows.StructScan(&category)
		if err != nil {
			return categories, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
		}

		categories = append(categories, category)
	}

	c.log.Debug(ctx, fmt.Sprintf("success to get all category with param: %v", param))

	return categories, nil
}
