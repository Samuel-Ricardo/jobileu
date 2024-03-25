package handler

import (
	"net/http"

	"github.com/Samuel-Ricardo/jobileu/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("Incorrect request body: %v", err.Error())
		sendError(ctx, http.StatusUnprocessableEntity, "Incorrect request body")
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation Error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "Error validating request")
	}

	opening := schemas.Openig{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error creating opening")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
