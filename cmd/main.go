package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/tokzy/eccom-rest-api/internal/env"
)

func main() {
	ctx := context.Background()
	cfg := config{
		addr: ":8090",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", ""),
		},
	}

	conn, err := pgx.Connect(ctx, cfg.db.dsn)

	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	log.Printf("Connection to db Succesful")

	api := application{
		config: cfg,
		db:     conn,
	}
	if err := api.run(api.mount()); err != nil {
		log.Printf("server fail to start, err: %s", err)
		os.Exit(1)
	}
}
