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
