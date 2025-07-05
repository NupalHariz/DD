package rest

import (
	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
)

// @Summary Create Assignment
// @Description Set a new assignment
// @Tags Assignment
// @Security BearerAuth
// @Param data body dto.CreateAssignmentParam true "assignment"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 404 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/assignments [POST]
func (r *rest) CreateAssignment(ctx *gin.Context) {
	var param dto.CreateAssignmentParam

	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	err := r.uc.Assignment.Create(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeCreated, nil, nil)
}

// @Description Update Assignment
// @Tags Assignment
// @Security BearerAuth
// @Param id path string true "assignment id"
// @Param data body dto.UpdateAssignmentParam true "assignment"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 404 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/assignments/{id} [PUT]
func (r *rest) UpdateAssignment(ctx *gin.Context) {
	var param dto.UpdateAssignmentParam

	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	if err := r.BindUri(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	err := r.uc.Assignment.Update(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, nil, nil)
}

// @Summary Get All Assignment
// @Description Get all assignment
// @Tags Assignment
// @Security BearerAuth
// @Produce json
// @Param category_id query string false "category id"
// @Param status query string false "status"
// @Param priority query string false "priority"
// @Success 200 {object} entity.HTTPResp{data=[]dto.GetAllAssignmentResponse}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 404 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/assignments [GET]
func (r *rest) GetAllAssignment(ctx *gin.Context) {
	var param dto.GetAllAssignmentParam

	if err := r.BindQuery(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	datas, err := r.uc.Assignment.GetAll(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, datas, nil)
}
