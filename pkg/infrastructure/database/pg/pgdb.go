package pg

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Pgdb *gorm.DB

func PgInit() {
	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()
	dsn := "host=localhost user=user password=password dbname=moaForm sslmode=disable"
	pdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("db connected ...")
	Pgdb = pdb
}

// you need to create your database first. manually
