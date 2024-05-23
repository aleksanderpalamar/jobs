package handler

import (
	"net/http"

	"github.com/aleksanderpalamar/jobs/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List Openings
// @Description List all job Openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(ctx *gin.Context) {
	opening := []schemas.Opening{}
	if err := db.Find(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "listOpening", opening)
}
