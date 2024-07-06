package repositories

import (
	"context"
	"moaformbuilder/internal/models"
)

type StaticFormRepository interface {
	Create(context.Context, models.StaticForm) (string, error)
	Update(context.Context, models.StaticForm) error
	Delete(context.Context, string) error
	GetByID(context.Context, string) error
	Get(context.Context) ([]models.StaticForm, error)
}

type StaticformFieldSelectionRepository interface {
	Create(context.Context, models.FormFieldSelection) (string, error)
	Update(context.Context, models.FormFieldSelection) error
	Delete(context.Context, string) error
	GetByID(context.Context, string) error
	Get(context.Context) ([]models.FormFieldSelection, error)
}
