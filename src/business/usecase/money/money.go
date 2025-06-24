package money

import (
	"context"
	"sync"

	budgetDom "github.com/NupalHariz/DD/src/business/domain/budget"
	moneyDom "github.com/NupalHariz/DD/src/business/domain/money"

	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/auth"
)

type Interface interface {
	Create(ctx context.Context, param dto.CreateTransactionParam) error
	Update(ctx context.Context, param dto.UpdateTransactionParam) error
}

type money struct {
	auth      auth.Interface
	moneyDom  moneyDom.Interface
	budgetDom budgetDom.Interface
}

type InitParam struct {
	Auth      auth.Interface
	MoneyDom  moneyDom.Interface
	BudgetDom budgetDom.Interface
}

func Init(param InitParam) Interface {
	return &money{
		auth:      param.Auth,
		moneyDom:  param.MoneyDom,
		budgetDom: param.BudgetDom,
	}
}

func (m *money) Create(ctx context.Context, param dto.CreateTransactionParam) error {
	loginUser, err := m.auth.GetUserAuthInfo(ctx)
	if err != nil {
		return err
	}

	// To Do: check if category exist

	inputMoneyParam := param.ToInputMoneyParam(loginUser.ID)

	var wg sync.WaitGroup
	errCh := make(chan error, 2)

	wg.Add(2)
	go func() {
		defer wg.Done()
		errCh <- m.budgetDom.UpdateExpense(
			ctx,
			entity.BudgetUpdateParam{
				UserId:         inputMoneyParam.UserId,
				CategoryId:     inputMoneyParam.CategoryId,
				CurrentExpense: inputMoneyParam.Amount,
			},
		)
	}()

	go func() {
		defer wg.Done()
		errCh <- m.moneyDom.Create(ctx, inputMoneyParam)
	}()

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *money) Update(ctx context.Context, param dto.UpdateTransactionParam) error {
	//Update Budget
	if param.Amount != 0 {
		money, err := m.moneyDom.Get(ctx, entity.MoneyParam{Id: param.Id})
		if err != nil {
			return err
		}

		amountChange := param.Amount - money.Amount

		err = m.budgetDom.UpdateExpense(ctx, entity.BudgetUpdateParam{
			UserId:         money.UserId,
			CategoryId:     money.CategoryId,
			CurrentExpense: amountChange,
		})

		if err != nil {
			return err
		}
	}

	//Update Money
	updateMoneyParam := param.ToMoneyUpdateParam()

	err := m.moneyDom.Update(ctx, updateMoneyParam, entity.MoneyParam{Id: param.Id})
	if err != nil {
		return err
	}

	return nil
}
