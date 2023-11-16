package app

import (
	"context"
	"log"
	"stock-services/internal/controller"
	"stock-services/internal/repo"
	"stock-services/pkg/postgres"
)

func Run() {
	log.Println("Stock-services start work...")
	ctx := context.Background()

	pg := postgres.NewPostgres(ctx)
	dbRepo := repo.NewRepository(pg)

	controller.NewRoute(ctx, dbRepo)
}
