package validator

type RegisterForm struct {
	UserName string `from:"username"	valid:"Required"`
	Password string `from:"password" 	valid:"Required"`
}
