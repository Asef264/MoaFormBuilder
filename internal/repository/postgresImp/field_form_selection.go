package postgresImp

import (
	"context"
	"fmt"
	"moaformbuilder/internal/models"
	repositories "moaformbuilder/internal/repository"
	"moaformbuilder/pkg/infrastructure/database/pg"
)

type FormFieldSelection struct{}

func NewRepoFormFieldSelection() repositories.StaticformFieldSelectionRepository {
	return &FormFieldSelection{}
}

func (m *FormFieldSelection) Create(ctx context.Context, formSelection models.FormFieldSelection) (string, error) {
	result := pg.Pgdb.Create(&formSelection)
	if result.Error != nil {
		return "", result.Error
	}
	return fmt.Sprintf("%d", formSelection.ID), nil
}

func (m *FormFieldSelection) Update(ctx context.Context, form models.FormFieldSelection) error {
	panic("not implemented")
}

func (m *FormFieldSelection) Delete(ctx context.Context, id string) error {
	panic("not implemented")
}

func (m *FormFieldSelection) GetByID(ctx context.Context, id string) error {
	panic("not implemented")
}

func (m *FormFieldSelection) Get(ctx context.Context) ([]models.FormFieldSelection, error) {
	panic("not implemented")
}
