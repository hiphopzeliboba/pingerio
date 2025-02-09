package router

import (
	"github.com/go-chi/chi"
	"net/http"
	"pingerio/backend/internal/api/handler"
)

type Router struct {
	containerHandler *handler.ContainerHandler
}

func NewRouter(containerHandler *handler.ContainerHandler) *Router {
	return &Router{containerHandler: containerHandler}
}

func (r *Router) Setup() http.Handler {
	router := chi.NewRouter()

	router.Get("/containers", r.containerHandler.GetContainers)
	router.Post("/containers", r.containerHandler.SaveContainers)

	router.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return router
}
