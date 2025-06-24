package category

import (
	"context"

	domCategory "github.com/NupalHariz/DD/src/business/domain/category"

	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/reyhanmichiels/go-pkg/v2/auth"
)

type Interface interface {
	Create(ctx context.Context, param dto.CreateCategoryParam) error
}

type category struct {
	categoryDom domCategory.Interface
	auth        auth.Interface
}

type InitParam struct {
	CategoryDom domCategory.Interface
	Auth        auth.Interface
}

func Init(param InitParam) Interface {
	return &category{
		categoryDom: param.CategoryDom,
		auth:        param.Auth,
	}
}

func (c *category) Create(ctx context.Context, param dto.CreateCategoryParam) error {
	loginUser, err := c.auth.GetUserAuthInfo(ctx)
	if err != nil {
		return err
	}

	categoryInputParam := param.ToCategoryInputParam(loginUser.ID)

	err = c.categoryDom.Create(ctx, categoryInputParam)
	if err != nil {
		return err
	}

	return nil
}
