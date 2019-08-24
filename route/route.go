package route

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"json2struct/web/controllers"
)

func InitRouter(app *iris.Application) {

	mvc.New(app.Party("/api")).Handle(controllers.NewTestController())

}
