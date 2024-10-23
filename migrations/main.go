package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
    // We provide the path to the migration files and the PostgreSQL connection information.
    m, err := migrate.New(
        "file://./", 
        "postgres://postgres:postgres@localhost:6432/productapp?sslmode=disable") 
    if err != nil {
        log.Fatalf("Migration start error: %v", err)
    }

    // The `Up` method is used to run migrations. This runs all migration files.
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatalf("Migration error: %v", err)
    }

    log.Println("Migration completed succesfully!")
}