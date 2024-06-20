package resolver

import (
	"post-api/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	serv *service.Service
}

func NewResolver(serv *service.Service) *Resolver {
	return &Resolver{serv: serv}
}
