package domain

import (
	"context"
	"moaformbuilder/internal/models"
	"moaformbuilder/internal/repository/postgresImp"
	services "moaformbuilder/internal/service"
)

type StaticForm struct{}

func NewStaticFormService() services.StaticFormService {
	return StaticForm{}
}

func (sf StaticForm) CreateStaticForm(ctx context.Context, req models.StaticForm) (string, error) {
	staticFormRepo := postgresImp.NewRepoStaticForm()
	id, err := staticFormRepo.Create(ctx, req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (sf StaticForm) UpdateStaticForm(context.Context, models.StaticForm) error {
	panic("not implemented")
}

func (sf StaticForm) DeleteStaticForm(context.Context, string) error {
	panic("not implemented")
}

func (sf StaticForm) GetStaticFormByID(context.Context, string) error {
	panic("not implemented")
}

func (sf StaticForm) GetStaticForms(context.Context) ([]models.StaticForm, error) {
	panic("not implemented")
}
