package assignment

import (
	"context"

	assignmentDom "github.com/NupalHariz/DD/src/business/domain/assignment"
	assignmentCategoryDom "github.com/NupalHariz/DD/src/business/domain/assignment_category"
	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/auth"
	"github.com/reyhanmichiels/go-pkg/v2/query"
)

type Interface interface {
	Create(ctx context.Context, param dto.CreateAssignmentParam) error
	Update(ctx context.Context, param dto.UpdateAssignmentParam) error
	GetAll(ctx context.Context, param dto.GetAllAssignmentParam) ([]dto.GetAllAssignmentResponse, error)
}

type assignment struct {
	auth          auth.Interface
	assignmentDom assignmentDom.Interface
	assignmentCategoryDom   assignmentCategoryDom.Interface
}

type InitParam struct {
	Auth        auth.Interface
	Assignment  assignmentDom.Interface
	AssignmentCategoryDom assignmentCategoryDom.Interface
}

func Init(param InitParam) Interface {
	return &assignment{
		auth:          param.Auth,
		assignmentDom: param.Assignment,
		assignmentCategoryDom:   param.AssignmentCategoryDom,
	}
}

func (a *assignment) Create(ctx context.Context, param dto.CreateAssignmentParam) error {
	loginUser, err := a.auth.GetUserAuthInfo(ctx)
	if err != nil {
		return err
	}

	assignmentInputParam := param.ToAssignmentInputParam(loginUser.ID)

	err = a.assignmentDom.Create(ctx, assignmentInputParam)
	if err != nil {
		return err
	}

	return nil
}

func (a *assignment) Update(ctx context.Context, param dto.UpdateAssignmentParam) error {
	assignmentUpdateParam := param.ToAssignmentUpdateParam()

	err := a.assignmentDom.Update(ctx, assignmentUpdateParam, entity.AssignmentParam{Id: param.Id})
	if err != nil {
		return err
	}

	return nil
}

func (a *assignment) GetAll(ctx context.Context, param dto.GetAllAssignmentParam) ([]dto.GetAllAssignmentResponse, error) {
	var datas []dto.GetAllAssignmentResponse

	loginUser, err := a.auth.GetUserAuthInfo(ctx)
	if err != nil {
		return datas, err
	}

	assignments, err := a.assignmentDom.GetAll(
		ctx,
		entity.AssignmentParam{
			UserId: loginUser.ID,
			Option: query.Option{
				DisableLimit: false,
			},
			PaginationParam: param.PaginationParam,
		})
	if err != nil {
		return datas, err
	}

	categoryIdMapSet := make(map[int64]struct{})
	for _, a := range assignments {
		categoryIdMapSet[a.CategoryId] = struct{}{}
	}

	var categoryIds []int64
	for id := range categoryIdMapSet {
		categoryIds = append(categoryIds, id)
	}

	categories, err := a.assignmentCategoryDom.GetAll(ctx, entity.AssignmentCategoryParam{Ids: categoryIds})
	if err != nil {
		return datas, err
	}

	categoryMap := make(map[int64]string)
	for _, c := range categories {
		categoryMap[c.Id] = c.Name
	}

	for _, a := range assignments {
		categoryName := categoryMap[a.CategoryId]

		data := dto.GetAllAssignmentResponse{
			Id: a.Id,
			Category: categoryName,
			Name: a.Name,
			Deadline: a.Deadline,
			Status: string(a.Status),
			Priority: string(a.Priority),
		}

		datas = append(datas, data)
	}

	return datas, nil
}
