package model

type User struct {
	Name     string `json:"name"`
	Password string `json: "password"`
}
type UserGetBYId struct{
	Id string `json:"id"`
}
