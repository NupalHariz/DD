package category

import (
	"context"
	"fmt"
	"strings"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
	"github.com/reyhanmichiels/go-pkg/v2/errors"
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
