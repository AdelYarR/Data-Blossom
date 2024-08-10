package api

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/AdelYarR/Data-Blossom/pkg/repository"
	"github.com/gorilla/mux"
)

type APIServer struct {
	config *Config
	router *mux.Router
	logger *slog.Logger
	db     *repository.PGRepo
}

func New(cfg *Config, db *repository.PGRepo) *APIServer {
	return &APIServer{
		config: cfg,
		router: mux.NewRouter(),
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		db: db,
	}
}

func (s *APIServer) Start() error {
	s.Handle()

	err := http.ListenAndServe(s.config.BindAddr, s.router)
	if err != nil {
		return err
	}

	return nil
}

func (s *APIServer) Handle() {
	s.router.HandleFunc("/api/lang", s.languageHandler).Methods(http.MethodGet)

	s.router.Use(s.middleware)
}
