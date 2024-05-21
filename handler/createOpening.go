package handler

import (
	"net/http"

	"github.com/aleksanderpalamar/jobs/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.Bind(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:        request.Role,
		Level:       request.Level,
		Company:     request.Company,
		Location:    request.Location,
		Remote:      *request.Remote,
		Link:        request.Link,
		Salary:      request.Salary,
		Title:       request.Title,
		Description: request.Description,
		Keywords:    request.Keywords,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "createOpening", opening)
}
