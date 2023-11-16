package controller

import (
	"context"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"log"
	"net/http"
	"stock-services/internal/repo"
)

type Controller struct {
	ctx context.Context
	db  repo.Storage
}

func NewController(ctx context.Context, db repo.Storage) *Controller {
	return &Controller{ctx, db}
}

func NewRoute(ctx context.Context, db repo.Storage) {
	r := rpc.NewServer()
	log.Println("RPC server start work...")

	r.RegisterCodec(json.NewCodec(), "application/json")
	err := r.RegisterService(NewController(ctx, db), "")
	if err != nil {
		log.Fatal("Error, register new service: ", err)
	}

	http.Handle("/rpc", r)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error, start new service on 8080: ", err)
	}
}
