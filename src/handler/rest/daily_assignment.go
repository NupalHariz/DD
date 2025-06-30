package rest

import (
	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
)

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

func (r *rest) GetAllDailyAssignment(ctx *gin.Context) {
	data, err := r.uc.DailyAssignment.GetAll(ctx.Request.Context())
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, data, nil)
}
