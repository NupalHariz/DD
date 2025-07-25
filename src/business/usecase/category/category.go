package category

import (
	"context"

	budgetDom "github.com/NupalHariz/DD/src/business/domain/budget"
	categoryDom "github.com/NupalHariz/DD/src/business/domain/category"
	"github.com/NupalHariz/DD/src/business/entity"

	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/reyhanmichiels/go-pkg/v2/auth"
)

type Interface interface {
	Create(ctx context.Context, param dto.CreateCategoryParam) error
	GetAll(ctx context.Context) ([]dto.GetAllCategoryResponse, error)
}

type category struct {
	categoryDom categoryDom.Interface
	budgetDom   budgetDom.Interface
	auth        auth.Interface
}

type InitParam struct {
	CategoryDom categoryDom.Interface
	BudgetDom   budgetDom.Interface
	Auth        auth.Interface
}

func Init(param InitParam) Interface {
	return &category{
		categoryDom: param.CategoryDom,
		budgetDom:   param.BudgetDom,
		auth:        param.Auth,
	}
}

func (c *category) Create(ctx context.Context, param dto.CreateCategoryParam) error {
	loginUser, err := c.auth.GetUserAuthInfo(ctx)
	if err != nil {
		return err
	}

	categoryInputParam := param.ToCategoryInputParam(loginUser.ID)

	category, err := c.categoryDom.Create(ctx, categoryInputParam)
	if err != nil {
		return err
	}

	err = c.budgetDom.Create(
		ctx,
		entity.BudgetInputParam{
			UserId:     loginUser.ID,
			CategoryId: category.Id,
			Amount:     0,
			Type:       entity.Weekly,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *category) GetAll(ctx context.Context) ([]dto.GetAllCategoryResponse, error) {
	var datas []dto.GetAllCategoryResponse

	loginUser, err := c.auth.GetUserAuthInfo(ctx)
	if err != nil {
		return datas, err
	}

	categories, err := c.categoryDom.GetAll(
		ctx,
		entity.CategoryParam{
			UserId: loginUser.ID,
		},
	)
	if err != nil {
		return datas, err
	}

	for _, c := range categories {
		data := dto.GetAllCategoryResponse{
			Id:   c.Id,
			Name: c.Name,
		}

		datas = append(datas, data)
	}

	return datas, nil
}
