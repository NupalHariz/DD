package rest

import (
	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
)

// @Summary Create Transaction
// @Description Set a new transaction
// @Tags Transaction
// @Security BearerAuth
// @Param data body dto.CreateTransactionParam true "transaction"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 404 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/moneys [POST]
func (r *rest) AddTransaction(ctx *gin.Context) {
	var param dto.CreateTransactionParam

	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	err := r.uc.Money.Create(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeCreated, nil, nil)
}

// @Summary Update Transaction
// @Description Update Transaction
// @Tags Transaction
// @Security BearerAuth
// @Param id path string true "money id"
// @Param data body dto.UpdateTransactionParam true "money"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 404 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/moneys/{id} [PUT]
func (r *rest) UpdateTransaction(ctx *gin.Context) {
	var param dto.UpdateTransactionParam

	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	if err := r.BindUri(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	err := r.uc.Money.Update(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, nil, nil)
}

// @Summary Get All Transaction
// @Description Get all transaction
// @Tags Transaction
// @Security BearerAuth
// @Produce json
// @Param category_id query string false "category id"
// @Param type query string false "type"
// @Success 200 {object} entity.HTTPResp{data=[]dto.GetTransactionResponse}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 404 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/moneys [GET]
func (r *rest) GetTransaction(ctx *gin.Context) {
	var param dto.GetTransactionParam

	if err := r.BindQuery(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	data, err := r.uc.Money.GetTransaction(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, data, nil)
}
