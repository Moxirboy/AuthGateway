package repo

import (
	"Auth_service-and-Api_gateway/proto"
	"Auth_service-and-Api_gateway/service/conn"
	_const "Auth_service-and-Api_gateway/service/const"
	"Auth_service-and-Api_gateway/service/model"
	"fmt"
)

type ty interface {
	*proto.UserResponse
}

func (auth *AuthPostgres) Create(u *proto.UserRequest) (res *proto.UserResponse, err error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (name ,username,password_hash) values($1,$2,$3) RETURNING id", _const.UsersTable)
	row := auth.Db.QueryRow(query, u.Name, u.Username, u.Password)
	if err := row.Scan(&id); err != nil {
		return nil, err
	}
	res.Id = id
	return res, nil
}
func GetUser(username string, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("select id from %s where username=$1 and password=$2")
	err := conn.Db.Get(&user, query)
	if err != nil {
		panic(err)
	}
	return user, nil
}
