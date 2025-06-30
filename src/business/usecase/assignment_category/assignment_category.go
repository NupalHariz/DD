package assignmentcategory

import (
	"context"

	assignmentCategoryDom "github.com/NupalHariz/DD/src/business/domain/assignment_category"
	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/auth"
)

type Interface interface {
	Create(ctx context.Context, param dto.CreateAssignmentCategory) error
	GetAll(ctx context.Context) ([]dto.GetAllAssignmentCategoryResponse, error)
}

type assignmentCategory struct {
	auth                  auth.Interface
	assignmentCategoryDom assignmentCategoryDom.Interface
}

type InitParam struct {
	Auth                  auth.Interface
	AssignmentCategoryDom assignmentCategoryDom.Interface
}

func Init(param InitParam) Interface {
	return &assignmentCategory{
		auth:                  param.Auth,
		assignmentCategoryDom: param.AssignmentCategoryDom,
	}
}

func (a *assignmentCategory) Create(ctx context.Context, param dto.CreateAssignmentCategory) error {
	loginUser, err := a.auth.GetUserAuthInfo(ctx)
	if err != nil {
		return err
	}

	assignmentCategoryInputParam := param.ToAssignmentCategoryInputParam(loginUser.ID)

	err = a.assignmentCategoryDom.Create(ctx, assignmentCategoryInputParam)
	if err != nil {
		return err
	}

	return nil
}

func (a *assignmentCategory) GetAll(ctx context.Context) ([]dto.GetAllAssignmentCategoryResponse, error) {
	var datas []dto.GetAllAssignmentCategoryResponse

	loginUser, err := a.auth.GetUserAuthInfo(ctx)
	if err != nil {
		return datas, err
	}

	assignmentCategories, err := a.assignmentCategoryDom.GetAll(
		ctx, 
		entity.AssignmentCategoryParam{
			UserId: loginUser.ID,
		},
	)
	if err != nil {
		return datas, err
	}

	for _, a := range assignmentCategories {
		data := dto.GetAllAssignmentCategoryResponse{
			Id: a.Id,
			Name: a.Name,
		}

		datas = append(datas, data)
	}

	return datas, nil
}
