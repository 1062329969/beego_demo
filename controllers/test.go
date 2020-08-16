package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Test_get() {
	c.Ctx.Output.Body([]byte("22222222222"))
	c.Ctx.Output.Body([]byte("hello world"))
}
func (c *TestController) Test_post() {
	c.Ctx.Output.Body([]byte("qqqqqqq"))
	c.Ctx.Output.Body([]byte("wwwwwwwww"))
}
func (c *TestController) Test2() {
	c.Ctx.WriteString("333333333333")
	c.Ctx.WriteString("55555555555555")
}

func (c *TestController) Anytest() {
	c.Ctx.WriteString("yyyyyyyyyyyyy")
}
