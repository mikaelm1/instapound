package main

import (
	"database/sql"
	"fmt"
	"os"

	"instapound/pkg/api"
	"instapound/pkg/repository"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"

	"instapound/utility"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	config, err := utility.NewConfig(utility.DEV)
	db, err := setupDB(config.DBDriver, config.DBSource)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	appStorage := repository.MakeAppStorage(db)
	err = appStorage.RunMigrations(config.DBSource, config.MigrationPath)
	if err != nil {
		return err
	}

	router := chi.NewRouter()
	server := api.NewServer(router)
	err = server.Run()
	if err != nil {
		return err
	}
	return nil
}

func setupDB(dbDriver string, dbSource string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		fmt.Printf("Error opening connection to DB: %v", err)
		return nil, err
	}

	// ping to make sure it's connected
	err = db.Ping()
	if err != nil {
		fmt.Printf("Failed to ping DB: %v", err)
		return nil, err
	}

	return db, nil
}
