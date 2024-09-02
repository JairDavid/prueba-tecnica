package api

import (
	"github.com/go-chi/chi/v5"
	"omnicloud.mx/tasks/pkg/infra/api/router"
)

type Server struct {
	engine     *chi.Mux
	taskRouter router.ITaskRouter
}

func New(engine *chi.Mux, taskRouter router.ITaskRouter) *Server {
	return &Server{
		engine:     engine,
		taskRouter: taskRouter,
	}
}

func (s *Server) MountRoutes() {
	s.engine.Mount("/api/v1", s.taskRouter.Mount())
}

func (s *Server) GetEngine() *chi.Mux {
	return s.engine
}
