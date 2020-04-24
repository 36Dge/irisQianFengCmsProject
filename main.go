package main

import (
	"github.com/kataras/iris"
	//"github.com/kataras/iris/_examples/mvc/login/datasource"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	config2 "irisDemo/QianFengCmsProject/config"
	"irisDemo/QianFengCmsProject/controller"
	"irisDemo/QianFengCmsProject/service"
	"time"
	"irisDemo/QianFengCmsProject/datasource"
)

func main() {

	app := newApp()

	// 应用App设置
	configuation(app)

	// 路由设置
	mvcHandle(app)

	config := config2.InitConfig()
	addr := ":" + config.Port
	app.Run(
		// 在端口9000进行监听
		iris.Addr(addr),
		// 无服务错误提示
		iris.WithoutServerError(iris.ErrServerClosed),
		//对json数据序列化更块的配置
		iris.WithOptimizations,
	)
}

// 构建app

func newApp() *iris.Application {
	app := iris.New()

	// 设置日志级别 开发阶段为debug
	app.Logger().SetLevel("debug")

	// 注册静态资源
	app.StaticWeb("/static", "./static")
	app.StaticWeb("/manager/static", "./static")

	// 注册视图文件
	app.RegisterView(iris.HTML("./static", ".html"))
	app.Get("/", func(context context.Context) {
		context.View("index,html")
	})
	return app
}

// 项目配置
func configuation(app *iris.Application) {

	// 配置 字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	// 错误配置
	// 未发现配置
	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusNotFound,
			"msg":    "not found",
			"data":   iris.Map{}, // 传空值
		})
	})

	// 服务器内部错误
	app.OnErrorCode(iris.StatusInternalServerError, func(context context.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusInternalServerError,
			"msg":    "interal error",
			"data":   iris.Map{},
		})
	})

}

// mvc架构处理

func mvcHandle(app *iris.Application) {
	// 启用session
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncooke",
		Expires: 24 * time.Hour,
	})

	// 获取redis实例
	redis := datasource.NewRedis()
	//设置session的同步位置为redis
	sessManager.UseDatabase(redis)

	// a 构造一个数据库引擎
	engine := datasource.NewMysqlEngine()

	// 管理员模块功能
	// 将实例化的数据库引擎放到adminservice当中
	adminService := service.AdminService(engine)

	admin := mvc.New(app.Party("/admin"))
	admin.Register(
		adminService,
		sessManager.Start,
	)
	admin.Handle(new(controller.AdminController))

	// 用户功能模块
	userService := service.NewAdminService(engine)
	user := mvc.New(app.Party("/v1/users"))
	user.Register(
		userService,
		sessManager.Start,
	)
	user.Handle(new(controller.Usercontroller))

	// 统计功能模块
	statisService := service.NewStatisService(engine)
	statis := mvc.New(app.Party("/statis/{model}/{data}/"))
	statis.Register(
		statisService,
		sessManager.Start,
	)
	statis.Handle(new(controller.StatisController))
}

// 订单模块

// 项目设置
