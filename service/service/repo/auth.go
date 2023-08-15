package repo

import (
	"Auth_service-and-Api_gateway/proto"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	Db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{Db: db}
}

type Auth interface {
	Create(u *proto.UserRequest) (*proto.UserResponse, error)
}
type Repository struct {
	Auth
}
