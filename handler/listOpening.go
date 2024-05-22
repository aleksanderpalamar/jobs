package handler

import (
	"net/http"

	"github.com/aleksanderpalamar/jobs/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	opening := []schemas.Opening{}
	if err := db.Find(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "listOpening", opening)
}
