package main

import (
	"context"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"

	"app/config"
	"app/repository"
	"app/server"
	"app/service"
	"app/util"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config")
	}

	db, err := util.NewPgxPool(cfg.Db.Url, context.Background())
	if err != nil {
		log.Fatal("cannot connect to db")
	}
	defer db.Close()

	runDBMigration(cfg.Migration.Url, cfg.Db.Url)

	repo := repository.NewSqlRepository(db)

	svc := service.GetService(repo)

	svr := server.NewGinServer(svc, &cfg)

	svr.Run()
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal(err, "cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err, "failed to run migrate up")
	}

	log.Println("db migrated successfully")
}
