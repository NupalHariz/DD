package rest

import (
	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
)

// @Summary Create Daily Assignment
// @Description Set a new daily assignment
// @Tags Daily Assignment
// @Security BearerAuth
// @Param data body dto.CreateDailyAssignmentParam true "daily assignment"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 404 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/daily-assignments [POST]
func (r *rest) CreateDailyAssignment(ctx *gin.Context) {
	var param dto.CreateDailyAssignmentParam

	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	err := r.uc.DailyAssignment.Create(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeCreated, nil, nil)
}

// @Summary Update Daily Assignment
// @Description Update daily assignment
// @Tags Daily Assignment
// @Security BearerAuth
// @Param id path string true "daily assignment id"
// @Param data body dto.UpdateDailyAssignmentParam true "daily assignment"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 404 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/daily-assignments/{id} [PUT]
func (r *rest) UpdateDailyAssignment(ctx *gin.Context) {
	var param dto.UpdateDailyAssignmentParam

	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	if err := r.BindUri(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	err := r.uc.DailyAssignment.Update(ctx, param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, nil, nil)
}

// @Summary Get All Daily Assignment
// @Description Get all daily assignment
// @Tags Daily Assignment
// @Security BearerAuth
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=[]dto.GetAllDailyAssignmentResponse}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 404 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/daily-assignments [GET]
func (r *rest) GetAllDailyAssignment(ctx *gin.Context) {
	data, err := r.uc.DailyAssignment.GetAll(ctx.Request.Context())
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, data, nil)
}
