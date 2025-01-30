package routes

import (
	"feedback-service/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/feedback", handlers.SendFeedbackHandler)
}
