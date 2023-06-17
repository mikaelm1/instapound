package repository

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type AppStorage interface {
	RunMigrations(connectionString string, migrationPath string) error
}

type appStorage struct {
	db *sql.DB
}

func MakeAppStorage(db *sql.DB) AppStorage {
	return &appStorage{db: db}
}

func (s *appStorage) RunMigrations(dbSource string, migrationPath string) error {
	m, err := migrate.New(migrationPath, dbSource)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Running migrations...")

	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			fmt.Println("No new migrations to apply")
		} else {
			return err
		}
	}

	version, _, err := m.Version()
	if err != nil {
		fmt.Println("No migrations have been applied yet")
	} else {
		fmt.Println("Current migration version:", version)
	}

	return nil
}
