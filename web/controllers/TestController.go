package controllers

import (
	"github.com/kataras/iris"
	"json2struct/service"
)

type TestController struct {
	Ctx         iris.Context
	TestService *service.TestService
	Common
}

func NewTestController() *TestController {

	return &TestController{TestService: service.NewTestService()}
}

func (t *TestController) Get() {
	t.ReturnSuccess()
}

func (t *TestController) PostRotate() {
	m := make(map[string]interface{})
	t.Ctx.ReadJSON(&m)
	userAgent := t.Ctx.Request().UserAgent()
	host := t.Ctx.Host()
	s, ok := t.TestService.ParseJsonStr(m, userAgent, host, "")
	if ok {
		t.ReturnSuccess(s)
	} else {
		t.ReturnJson(10001, "转换失败")
	}
}

func (t *TestController) GetRotateCount() {
	var count int
	t.TestService.GetRotateCount(&count)
	t.ReturnSuccess(count)
}
