package handler

import (
	"net/http"

	"github.com/Samuel-Ricardo/jobileu/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Openig{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
