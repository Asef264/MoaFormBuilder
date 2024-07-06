package postgresImp

import (
	"context"
	"fmt"
	"moaformbuilder/internal/models"
	repositories "moaformbuilder/internal/repository"
	"moaformbuilder/pkg/infrastructure/database/pg"
)

type StaticForm struct{}

func NewRepoStaticForm() repositories.StaticFormRepository {
	return &StaticForm{}
}

func (m *StaticForm) Create(ctx context.Context, form models.StaticForm) (string, error) {
	result := pg.Pgdb.Create(&form)
	if result.Error != nil {
		return "", result.Error
	}
	return fmt.Sprintf("%d", form.ID), nil
}

func (m *StaticForm) Update(ctx context.Context, form models.StaticForm) error {
	panic("not implemented")
}

func (m *StaticForm) Delete(ctx context.Context, id string) error {
	panic("not implemented")
}

func (m *StaticForm) GetByID(ctx context.Context, id string) error {
	panic("not implemented")
}

func (m *StaticForm) Get(ctx context.Context) ([]models.StaticForm, error) {
	panic("not implemented")
}
