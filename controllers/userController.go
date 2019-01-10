package controllers

import (
	"github.com/astaxie/beego"
	"goblog/models"
	"goblog/validator"
	"time"
)

type UserController struct {
	BaseController
}

func (c *UserController) Register() {
	form := validator.RegisterForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseRegisterForm:", err)
		c.Data["json"] = validator.NewErrorInfo(ErrInputData)
		c.ServeJSON()
		return
	}

	beego.Debug("ParseRegisterForm : ", &form)
	if err := c.VerifyForm(&form); err != nil {
		beego.Debug("ValidateRegisterForm", err)
		c.Data["json"] = validator.NewErrorInfo(ErrInputData)
		c.ServeJSON()
		return
	}

	beego.Debug("validateRegisterForm", &form)
	regData := time.Now()
	user, err := models.NewUser(&form, regData)
	if err != nil {
		beego.Debug("NewUser:", err)
		c.Data["json"] = validator.NewErrorInfo(ErrSystem)
		c.ServeJSON()
		return
	}
	beego.Debug("NewUser", user)

	if code, err := user.Insert(); err != nil {
		beego.Error("InsertUser:", err)
		if code == models.ErrDupRows {
			c.Data["json"] = models.NewErrorInfo(ErrDupUser)
		} else {
			c.Data["json"] = models.NewErrorInfo(ErrDatabase)
		}
		c.ServeJSON()
		return
	}

	go models.IncTotalUserCount(regDate)

	c.Data["json"] = validator.NewNormalInfo("Succes")
	c.ServeJSON()

}
