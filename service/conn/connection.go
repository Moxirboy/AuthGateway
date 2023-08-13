package conn

import (
	"Auth_service-and-Api_gateway/service/config"
	"Auth_service-and-Api_gateway/service/pkg/postgres"
)

var Db, _ = postgres.Conn(config.Configs())
