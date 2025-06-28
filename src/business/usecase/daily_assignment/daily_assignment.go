package dailyassignment

import (
	"context"

	dailyAssignmentDom "github.com/NupalHariz/DD/src/business/domain/daily_assignment"
	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/reyhanmichiels/go-pkg/v2/auth"
)

type Interface interface {
	Create(ctx context.Context, param dto.CreateDailyAssignmentParam) error
}

type dailyAssignment struct {
	auth               auth.Interface
	dailyAssignmentDom dailyAssignmentDom.Interface
}

type InitParam struct {
	Auth               auth.Interface
	DailyAssignmentDom dailyAssignmentDom.Interface
}

func Init(param InitParam) Interface {
	return &dailyAssignment{
		auth:               param.Auth,
		dailyAssignmentDom: param.DailyAssignmentDom,
	}
}

func (d *dailyAssignment) Create(ctx context.Context, param dto.CreateDailyAssignmentParam) error {
	loginUser, err := d.auth.GetUserAuthInfo(ctx)
	if err != nil {
		return err
	}

	dailyAssignmentInputParam := param.ToDailyAssignmentInputParam(loginUser.ID)

	err = d.dailyAssignmentDom.Create(ctx, dailyAssignmentInputParam)
	if err != nil {
		return err
	}

	return nil
}
