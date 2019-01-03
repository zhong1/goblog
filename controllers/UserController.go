package controllers

type UserController struct {
	BaseController
}

func (c *UserController) Register() {
	c.Ctx.WriteString("hello index")
}


