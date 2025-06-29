package money

import (
	"context"
	"sync"

	budgetDom "github.com/NupalHariz/DD/src/business/domain/budget"
	categoryDom "github.com/NupalHariz/DD/src/business/domain/category"
	moneyDom "github.com/NupalHariz/DD/src/business/domain/money"

	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/auth"
	"github.com/reyhanmichiels/go-pkg/v2/null"
	"github.com/reyhanmichiels/go-pkg/v2/query"
)

type Interface interface {
	Create(ctx context.Context, param dto.CreateTransactionParam) error
	Update(ctx context.Context, param dto.UpdateTransactionParam) error
	GetTransaction(ctx context.Context, param dto.GetTransactionParam) ([]dto.GetTransactionResponse, error)
}

type money struct {
	auth        auth.Interface
	moneyDom    moneyDom.Interface
	budgetDom   budgetDom.Interface
	categoryDom categoryDom.Interface
}

type InitParam struct {
	Auth        auth.Interface
	MoneyDom    moneyDom.Interface
	BudgetDom   budgetDom.Interface
	CategoryDom categoryDom.Interface
}

func Init(param InitParam) Interface {
	return &money{
		auth:        param.Auth,
		moneyDom:    param.MoneyDom,
		budgetDom:   param.BudgetDom,
		categoryDom: param.CategoryDom,
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
				CurrentExpense: null.Int64From(inputMoneyParam.Amount),
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
	if param.Amount != 0 {
		money, err := m.moneyDom.Get(ctx, entity.MoneyParam{Id: param.Id})
		if err != nil {
			return err
		}

		amountChange := param.Amount - money.Amount

		err = m.budgetDom.UpdateExpense(ctx, entity.BudgetUpdateParam{
			UserId:         money.UserId,
			CategoryId:     money.CategoryId,
			CurrentExpense: null.Int64From(amountChange),
		})

		if err != nil {
			return err
		}
	}

	updateMoneyParam := param.ToMoneyUpdateParam()

	err := m.moneyDom.Update(ctx, updateMoneyParam, entity.MoneyParam{Id: param.Id})
	if err != nil {
		return err
	}

	return nil
}

func (m *money) GetTransaction(ctx context.Context, param dto.GetTransactionParam) ([]dto.GetTransactionResponse, error) {
	var res []dto.GetTransactionResponse

	loginUser, err := m.auth.GetUserAuthInfo(ctx)
	if err != nil {
		return res, err
	}

	moneys, err := m.moneyDom.GetAll(ctx,
		entity.MoneyParam{
			UserId:          loginUser.ID,
			CategoryId:      param.CategoryId,
			Type:            entity.MoneyType(param.Type),
			Option:          query.Option{DisableLimit: false},
			PaginationParam: param.PaginationParam,
		})

	if err != nil {
		return res, err
	}

	categoryIdSet := make(map[int64]struct{})
	for _, m := range moneys {
		categoryIdSet[m.CategoryId] = struct{}{}
	}

	var categoryIds []int64
	for id := range categoryIdSet {
		categoryIds = append(categoryIds, id)
	}

	categories, err := m.categoryDom.GetAll(ctx, entity.CategoryParam{Ids: categoryIds})
	if err != nil {
		return res, err
	}

	categoryMap := make(map[int64]string)
	for _, c := range categories {
		categoryMap[c.Id] = c.Name
	}

	for _, m := range moneys {
		categoryName := categoryMap[m.CategoryId]

		transactionResp := dto.GetTransactionResponse{
			Id:       m.Id,
			Amount:   m.Amount,
			Category: categoryName,
			Type:     m.Type,
		}

		res = append(res, transactionResp)
	}

	return res, nil
}
