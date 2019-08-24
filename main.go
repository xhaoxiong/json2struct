package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"json2struct/config"
	"json2struct/models"
	"json2struct/route"
)

var (
	cfg = pflag.StringP("config", "c", "", "./config.yaml")
)

func main() {
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	models.DB.Init()
	app := newApp()

	route.InitRouter(app)

	app.Run(iris.Addr(viper.GetString("addr")))
}

func newApp() *iris.Application {
	app := iris.New()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	app.Use(crs) //
	app.StaticWeb("/static", "./web/views/static")
	app.RegisterView(iris.HTML("./web/views/", ".html"))
	app.AllowMethods(iris.MethodOptions)
	//app.Use(middleware.GetJWT().Serve)//是否启用jwt中间件
	app.Configure(iris.WithOptimizations)
	app.Get("/", func(ctx context.Context) {
		ctx.View("index.html")
	})
	return app
}
