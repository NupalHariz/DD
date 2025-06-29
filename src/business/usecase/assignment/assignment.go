package assignment

import (
	"context"

	assignmentDom "github.com/NupalHariz/DD/src/business/domain/assignment"
	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/reyhanmichiels/go-pkg/v2/auth"
)

type Interface interface {
	Create(ctx context.Context, param dto.CreateAssignmentParam) error
}

type assignment struct {
	auth          auth.Interface
	assignmentDom assignmentDom.Interface
}

type InitParam struct {
	Auth       auth.Interface
	Assignment assignmentDom.Interface
}

func Init(param InitParam) Interface {
	return &assignment{
		auth:          param.Auth,
		assignmentDom: param.Assignment,
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
