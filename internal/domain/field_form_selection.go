package domain

import (
	"context"
	"moaformbuilder/internal/models"
	"moaformbuilder/internal/repository/postgresImp"
	services "moaformbuilder/internal/service"
)

type FormFieldSelection struct{}

func NewFormFieldSelectionService() services.FormFieldSelectionService {
	return FormFieldSelection{}
}

func (sf FormFieldSelection) CreateFormFieldSelection(ctx context.Context, req models.FormFieldSelection) (string, error) {
	FormFieldSelectionRepo := postgresImp.NewRepoFormFieldSelection()
	id, err := FormFieldSelectionRepo.Create(ctx, req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (sf FormFieldSelection) UpdateFormFieldSelection(context.Context, models.FormFieldSelection) error {
	panic("not implemented")
}

func (sf FormFieldSelection) DeleteFormFieldSelection(context.Context, string) error {
	panic("not implemented")
}

func (sf FormFieldSelection) GetFormFieldSelectionByID(context.Context, string) error {
	panic("not implemented")
}

func (sf FormFieldSelection) GetFormFieldSelections(context.Context) ([]models.FormFieldSelection, error) {
	panic("not implemented")
}
