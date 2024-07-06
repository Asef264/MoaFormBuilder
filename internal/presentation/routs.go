package presentation

import (
	"github.com/gin-gonic/gin"
)

func PresentationInit() {
	router := gin.New()
	formFieldSelection := router.Group("")
	formFieldSelection.POST(("/create"), POSTCreateFormFieldSelection)
	router.Run(":8080")
}
