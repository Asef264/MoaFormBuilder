package presentation

import (
	"log"
	"moaformbuilder/internal/domain"
	"moaformbuilder/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func POSTCreateFormFieldSelection(ctx *gin.Context) {
	var formFieldSelection models.FormFieldSelection
	if err := ctx.BindJSON(&formFieldSelection); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missmatch type"})
		return
	}
	formFieldSelectionSrv := domain.NewFormFieldSelectionService()
	id, err := formFieldSelectionSrv.CreateFormFieldSelection(ctx, formFieldSelection)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusAccepted, id)
}
