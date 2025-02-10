package database

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate() {
	m, err := migrate.New("file://db/migrations", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
	}
	if err := m.Up(); err != nil {
		fmt.Println(err)
	}
}
