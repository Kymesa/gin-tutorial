package transport

import "gin-tutorial/internal/services"

type BookHander struct {
	services *services.Services
}

func New(s *services.Services) *BookHander {
	return &BookHander{
		services: s,
	}
}
