package services

import (
	"context"
	"moaformbuilder/internal/models"
)

type FormFieldSelectionService interface {
	CreateFormFieldSelection(context.Context, models.FormFieldSelection) (string, error)
	UpdateFormFieldSelection(context.Context, models.FormFieldSelection) error
	DeleteFormFieldSelection(context.Context, string) error
	GetFormFieldSelectionByID(context.Context, string) error
	GetFormFieldSelections(context.Context) ([]models.FormFieldSelection, error)
}
