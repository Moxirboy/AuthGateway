package models

type UserCreate struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
