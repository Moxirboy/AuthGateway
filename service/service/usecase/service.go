package usecase

import (
	"Auth_service-and-Api_gateway/proto"
	"Auth_service-and-Api_gateway/service/pkg/jwt"
	"context"
)

func (s *Server) Create(ctx context.Context, req *proto.UserRequest) (res *proto.UserResponse, err error) {
	req.Password = jwt.GeneratePasswordHash(req)
	return s.repo.Create(req)
}
