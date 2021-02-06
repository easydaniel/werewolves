package models

type User struct {
	Base
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Password    string `json:"-"`
}
