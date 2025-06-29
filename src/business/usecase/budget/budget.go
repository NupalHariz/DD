package budget

import (
	"context"
	"time"

	budgetDom "github.com/NupalHariz/DD/src/business/domain/budget"
	categoryDom "github.com/NupalHariz/DD/src/business/domain/category"
	historyBudgetDom "github.com/NupalHariz/DD/src/business/domain/history_budget"
	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/auth"
	"github.com/reyhanmichiels/go-pkg/v2/null"
)

type Interface interface {
	Create(ctx context.Context, param dto.CreateBudgetParam) error
	Update(ctx context.Context, param dto.UpdateBudgetParam) error
	GetAll(ctx context.Context, param dto.GetBudgetParam) ([]dto.GetAllBudgetResponse, error)
	WeeklyResetScheduler(ctx context.Context) error
	MonthlyResetScheduler(ctx context.Context) error
}

type budget struct {
	auth             auth.Interface
	budgetDom        budgetDom.Interface
	historyBudgetDom historyBudgetDom.Interface
	categoryDom      categoryDom.Interface
}

type InitParam struct {
	Auth             auth.Interface
	BudgetDom        budgetDom.Interface
	HistoryBudgetDom historyBudgetDom.Interface
	CategoryDom      categoryDom.Interface
}

func Init(param InitParam) Interface {
	return &budget{
		auth:             param.Auth,
		budgetDom:        param.BudgetDom,
		historyBudgetDom: param.HistoryBudgetDom,
		categoryDom:      param.CategoryDom,
	}
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

func (b *budget) GetAll(ctx context.Context, param dto.GetBudgetParam) ([]dto.GetAllBudgetResponse, error) {
	var datas []dto.GetAllBudgetResponse

	loginUser, err := b.auth.GetUserAuthInfo(ctx)
	if err != nil {
		return datas, err
	}

	budgets, err := b.budgetDom.GetAll(
		ctx,
		entity.BudgetParam{
			UserId:          loginUser.ID,
			Type:            param.Type,
			PaginationParam: param.PaginationParam,
		},
	)
	if err != nil {
		return datas, err
	}

	categoryIdSet := make(map[int64]struct{})
	for _, b := range budgets {
		categoryIdSet[b.CategoryId] = struct{}{}
	}

	var categoryIds []int64
	for categoryId := range categoryIdSet {
		categoryIds = append(categoryIds, categoryId)
	}

	categories, err := b.categoryDom.GetAll(ctx, entity.CategoryParam{Ids: categoryIds})
	if err != nil {
		return datas, err
	}

	categoryMap := make(map[int64]string)
	for _, c := range categories {
		categoryMap[c.Id] = c.Name
	}

	for _, b := range budgets {
		categoryName := categoryMap[b.CategoryId]

		data := dto.GetAllBudgetResponse{
			Id:             b.Id,
			Category:       categoryName,
			Amount:         b.Amount,
			CurrentExpense: b.CurrentExpense,
			Type:           b.Type,
		}

		datas = append(datas, data)
	}

	return datas, nil
}

func (b *budget) WeeklyResetScheduler(ctx context.Context) error {
	return b.resetScheduler(ctx, entity.Weekly, b.getStartAndEndWeek)
}

func (b *budget) MonthlyResetScheduler(ctx context.Context) error {
	return b.resetScheduler(ctx, entity.Monthly, b.getStartAndEndMonth)
}

func (b *budget) resetScheduler(ctx context.Context, budgetType entity.BudgetType, startEndFn func(time.Time) (time.Time, time.Time)) error {
	now := time.Now()

	budgets, err := b.budgetDom.GetAll(ctx, entity.BudgetParam{Type: string(budgetType)})
	if err != nil {
		return err
	}

	start, end := startEndFn(now)

	var historyBudgets []entity.HistoryBudget

	var budgetParam entity.BudgetParam

	for _, b := range budgets {
		historyBudget := b.ToHistoryBudget(start, end)

		historyBudgets = append(historyBudgets, historyBudget)
		budgetParam.Ids = append(budgetParam.Ids, b.Id)
	}

	err = b.historyBudgetDom.CreateBatch(ctx, historyBudgets)
	if err != nil {
		return err
	}

	err = b.budgetDom.Update(
		ctx,
		entity.BudgetUpdateParam{
			CurrentExpense: null.Int64From(0),
		},
		budgetParam,
	)
	if err != nil {
		return err
	}

	return nil
}

func (b *budget) getStartAndEndWeek(now time.Time) (time.Time, time.Time) {
	weekday := now.Weekday()

	if weekday == 0 {
		weekday = 7
	}

	start := now.AddDate(0, 0, -int(weekday)+1)
	end := start.AddDate(0, 0, 6)

	return start, end
}

func (b *budget) getStartAndEndMonth(now time.Time) (time.Time, time.Time) {
	currentYear, currentMonth, _ := now.Date()

	start := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.Local)
	end := start.AddDate(0, 1, -1)

	return start, end
}
