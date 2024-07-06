package main

import (
	"moaformbuilder/internal/models"
	"moaformbuilder/pkg/infrastructure/database/pg"
)

func main() {
	serverInit()
}

func serverInit() {
	pg.PgInit()
	err := pg.Pgdb.AutoMigrate(
		&models.Register{},
		&models.FormFieldSelection{},
		&models.File{},
		&models.StaticForm{},
	)
	if err != nil {
		panic("failed on connecting db")
	}
}
