package services

import (
	"context"
	"moaformbuilder/internal/models"
)

type StaticFormService interface {
	CreateStaticForm(context.Context, models.StaticForm) (string, error)
	UpdateStaticForm(context.Context, models.StaticForm) error
	DeleteStaticForm(context.Context, string) error
	GetStaticFormByID(context.Context, string) error
	GetStaticForms(context.Context) ([]models.StaticForm, error)
}
