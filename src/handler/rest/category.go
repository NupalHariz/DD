package rest

import (
	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
)

// @Summary Create Category
// @Description Create category for users
// @Tags categories
// @Produce json
// @Param data body dto.CreateCategoryParam{} true "New Category"
// @Security BearerAuth
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 404 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/categories/ [POST]
func (r *rest) CreateCategory(ctx *gin.Context) {
	var param dto.CreateCategoryParam

	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	err := r.uc.Category.Create(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeCreated, nil, nil)
}

// @Summary Get All Category
// @Description Get all category that has been created by user
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Success 200 {object} entity.HTTPResp{data=[]dto.GetAllCategoryResponse}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 404 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/categories/ [GET]
func (r *rest) GetAllCategory(ctx *gin.Context) {
	data, err := r.uc.Category.GetAll(ctx.Request.Context())
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, data, nil)
}
