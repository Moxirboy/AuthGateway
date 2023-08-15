package usecase

import (
	"Auth_service-and-Api_gateway/proto"
	"Auth_service-and-Api_gateway/service/service/repo"
)

type Server struct {
	proto.AuthServer
	repo repo.Auth
}

func NewServer(repo repo.Auth) *Server {
	return &Server{repo: repo}
}
