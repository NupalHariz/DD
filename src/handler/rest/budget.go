package rest

import (
	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
)

func (r *rest) CreateBudget(ctx *gin.Context) {
	var param dto.CreateBudgetParam

	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	err := r.uc.Budget.Create(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeCreated, nil, nil)
}

// @Summary Update Budget
// @Description Set a new budget plan
// @Tags Budget
// @Security BearerAuth
// @Param id path string true "budget id"
// @Param data body dto.UpdateBudgetParam true "Update Budget Data"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 404 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/budgets/{id} [PUT]
func (r *rest) UpdateBudget(ctx *gin.Context) {
	var param dto.UpdateBudgetParam

	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	if err := r.BindUri(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	err := r.uc.Budget.Update(ctx, param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, nil, nil)
}

// @Summary Get All Budget
// @Description Get all budget
// @Tags Budget
// @Security BearerAuth
// @Param type query string flse "budget type"
// @Param page query string false "page"
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=[]dto.GetAllBudgetResponse}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 404 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/budgets [GET]
func (r *rest) GetAllBudget(ctx *gin.Context) {
	var param dto.GetBudgetParam

	if err := r.BindQuery(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	data, err := r.uc.Budget.GetAll(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, data, nil)
}
