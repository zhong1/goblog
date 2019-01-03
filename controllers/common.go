package controllers

import "github.com/astaxie/beego"

//控制器错误返回结构
type ControllerError struct {
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//公共控制器
type BaseController struct {
	beego.Controller
}

//错误返回
func (c *BaseController) RetError(e *ControllerError) {

}
