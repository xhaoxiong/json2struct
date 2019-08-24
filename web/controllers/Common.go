package controllers

import "github.com/kataras/iris"

type Common struct {
	Ctx iris.Context
}

func (c *Common) ReturnJson(status int, message string, args ...interface{}) {
	result := make(map[string]interface{})
	result["code"] = status
	result["message"] = message

	key := ""

	for _, arg := range args {
		switch arg.(type) {
		case string:
			key = arg.(string)

		default:
			result[key] = arg
		}
	}
	c.Ctx.JSON(result)
	c.Ctx.StopExecution()
	return
}

func (c *Common) ReturnSuccess(args ...interface{}) {
	result := make(map[string]interface{})
	result["code"] = 10000
	result["message"] = "success"
	for _, arg := range args {
		switch arg.(type) {
		default:
			result["data"] = arg
		}
	}
	c.Ctx.JSON(result)
	c.Ctx.StopExecution()
	return
}
