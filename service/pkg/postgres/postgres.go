package postgres

import (
	"Auth_service-and-Api_gateway/service/config"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func Conn(cfg *config.Config) (*sqlx.DB, error) {

	DbString := fmt.Sprintf("Port=%s,Host=%s,Password=%s,User=%s,DatabaseName=%s", cfg.User, cfg.Host, cfg.Password, cfg.User, cfg.Database)
	conn, err := sqlx.Open("pgx", DbString)
	if err != nil {
		panic(err)
	}

	return conn, nil
}
