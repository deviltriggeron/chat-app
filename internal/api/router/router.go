package router

import (
	"chat-app/internal/api/handler"

	"github.com/gorilla/mux"
)

type Router struct {
}

func NewRouter(handler handler.Handler) *mux.Router {
	m := mux.NewRouter()

	return m
}
