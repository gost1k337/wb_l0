package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gost1k337/wb_demo/subscriber/internal/service"
	"github.com/gost1k337/wb_demo/subscriber/pkg/logging"
	"github.com/rs/cors"
	"net/http"
	"strconv"
)

type Handler struct {
	services *service.Services
	http     *chi.Mux
	logger   logging.Logger
}

func New(services *service.Services, logger logging.Logger) *Handler {
	h := &Handler{
		services: services,
		logger:   logger,
	}

	corsCfg := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods: []string{
			http.MethodOptions,
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders: []string{"Accept", "Content-Type", "Accept-Encoding"},
	})

	r := chi.NewRouter()
	r.Use(corsCfg.Handler)
	r.Use(middleware.DefaultLogger)

	r.Get("/orders/{id}", h.GetOrder)

	h.http = r

	return h
}

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	order, err := h.services.GetById(id)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(&order); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *Handler) HTTP() http.Handler {
	return h.http
}
