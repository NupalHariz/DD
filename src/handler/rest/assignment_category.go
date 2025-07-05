package rest

import (
	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
)

// @Summary Create Assignment Category
// @Description Create a new assignment category
// @Tags Assignment Category
// @Security BearerAuth
// @Param data body dto.CreateAssignmentCategory true "New Assignment Category"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 404 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/assignment-categories [POST]
func (r *rest) CreateAssignmentCategory(ctx *gin.Context) {
	var param dto.CreateAssignmentCategory

	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	err := r.uc.AssignmentCategory.Create(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeCreated, nil, nil)
}

// @Summary Get All Assignment Category
// @Description Get all a new assignment category
// @Tags Assignment Category
// @Security BearerAuth
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=[]dto.GetAllAssignmentCategoryResponse}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 404 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/assignment-categories [GET]
func (r *rest) GetAllAssignmentCategory(ctx *gin.Context) {
	res, err := r.uc.AssignmentCategory.GetAll(ctx.Request.Context())
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}
