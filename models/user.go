package models

import (
	"beego-demo/models/mymysql"
	"crypto/rand"
	"fmt"
	"goblog/validator"
	"golang.org/x/crypto/scrypt"
	"io"
	"time"
)

type User struct {
	Id 	int
	UserName string `orm:"size(50)"`
	Password string `orm:"size(50)"`
	RegDate  time.Time `bson:"reg_date" json:"reg_date,omitempty"`
}

const pwHashBytes = 64

func generateSalt() (salt string, err error) {
	buf := make([]byte, pwHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", buf), nil
}

func generatePassHash(password string, salt string) (hash string, err error) {
	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, pwHashBytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h), nil
}

func NewUser(r *validator.RegisterForm, t time.Time) (u *User, err error)  {
	salt, err := generateSalt()
	if err != nil {
		return nil, err
	}
	hash, err := generatePassHash(r.Password, salt)
	if err != nil {
		return nil, err
	}

	user := User{
		UserName:     r.UserName,
		Password: hash,
		RegDate:  t}

	return &user, nil
}

// Insert insert a document to collection.
func (u *User) Insert() (code int, err error) {
	mConn := mymysql.Conn()
	defer mConn.Close()

	//c := mConn.DB("").C("users")
	//err = c.Insert(u)
	//
	//if err != nil {
	//	if mgo.IsDup(err) {
	//		code = ErrDupRows
	//	} else {
	//		code = ErrDatabase
	//	}
	//} else {
	//	code = 0
	//}
	return
}