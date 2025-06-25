package budget

import (
	"context"

	budgetDom "github.com/NupalHariz/DD/src/business/domain/budget"
	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/auth"
)

type Interface interface {
	Create(ctx context.Context, param dto.CreateBudgetParam) error
	Update(ctx context.Context, param dto.UpdateBudgetParam) error
}

type budget struct {
	auth      auth.Interface
	budgetDom budgetDom.Interface
}

type InitParam struct {
	Auth      auth.Interface
	BudgetDom budgetDom.Interface
}

func Init(param InitParam) Interface {
	return &budget{auth: param.Auth, budgetDom: param.BudgetDom}
}

func (b *budget) Create(ctx context.Context, param dto.CreateBudgetParam) error {
	loginUser, err := b.auth.GetUserAuthInfo(ctx)
	if err != nil {
		return err
	}

	inputBudgetParam := param.ToBudgetInputParam(loginUser.ID)

	err = b.budgetDom.Create(ctx, inputBudgetParam)
	if err != nil {
		return err
	}

	return nil
}

func (b *budget) Update(ctx context.Context, param dto.UpdateBudgetParam) error {
	budgetUpdateParam := param.ToBudgetUpdateParam()

	err := b.budgetDom.Update(ctx, budgetUpdateParam, entity.BudgetParam{Id: param.Id})
	if err != nil {
		return err
	}

	return nil
}
