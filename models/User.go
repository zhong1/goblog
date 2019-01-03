package models

type User struct {
	Id 	int
	UserName string `orm:"size(50)"`
	Password string `orm:"size(50)"`
}