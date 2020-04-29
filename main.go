package main

import (
	"github.com/kataras/iris"
	"io"
	"irisDemo/QianFengCmsProject/model"
	"irisDemo/QianFengCmsProject/utils"
	"os"
	"strconv"

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

	// 设定应用图标
	app.Favicon("./static/favicons/favicon.ico")
	// 设置日志级别 开发阶段为debug
	app.Logger().SetLevel("debug")

	// 注册静态资源
	app.StaticWeb("/static", "./static")
	app.StaticWeb("/manager/static", "./static")
	app.StaticWeb("/img", "./uploads")

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

	// 获取用户详细信息
	app.Get("/v1/user/{user_name}", func(context context.Context) {
		userName := context.Params().Get(user_name)
		var user model.User
		_, err := engine.Where("user_name= ?", userName).Get(&user)
		if err != nil {
			context.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERINFO,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ORDERINFO),
			})

		} else {
			context.JSON(user)
		}
	})

	// 获取地址信息
	app.Get("/v1/address/{address_id}", func(context context.Context) {
		address_id := context.Params().Get("address_id")
		addressID, err := strconv.Atoi(address_id)
		if err != nil {
			context.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERINFO,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ORDERINFO),
			})
		}
		var address model.Address
		_, err = engine.Id(addressID).Get(&address)
		if err != nil {
			context.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERINFO,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ORDERINFO),
			})
		}

		// 查询数据成功
		context.JSON(address)

	})

	// 统计功能模块
	statisService := service.NewStatisService(engine)
	statis := mvc.New(app.Party("/statis/{model}/{data}/"))
	statis.Register(
		statisService,
		sessManager.Start,
	)
	statis.Handle(new(controller.StatisController))
	// 订单模块
	orderService := service.NewOrderService(engine)
	order := mvc.New(app.Party("/bos/order/"))
	order.Register(
		orderService,
		sessManager.Start,
	)
	order.Handle(new(controller.OrderController)) // 控制器

	// 商铺模块
	shopService := service.NewShopService(engine)
	shop := mvc.New(app.Party("/shopping/restaurants/"))
	shop.Register(
		shopService,
		sessManager.Start,
		)
	shop.Handle(new(controller.ShopController)) // 控制器

	// 项目设置

	// 文件上传

	app.Post("/admin/update/avatar/{adminId}", func(context context.Context) {
		adminId := context.Params().Get("adminId")
		iris.New().Logger().Info(adminId)
		file, info, err := context.FormFile("file")
		if err != nil {
			iris.New().Logger().Info(err.Error())
			context.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_CATEGORYADD,
				"failure": utils.Recode2Text(utils.RESPMSG_ERROR_PICTUREADD),
			})
			return
		}
		defer file.Close()
		fname := info.Filename
		out, err := os.OpenFile("./uploads/"+fname, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			iris.New().Logger().Info(err.Error())
			context.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_CATEGORYADD,
				"failure": utils.Recode2Text(utils.RESPMSG_ERROR_PICTUREADD),
			})
			return
		}

		iris.New().Logger().Info("文件路径：" + out.Name())
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			context.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_CATEGORYADD,
				"failure": utils.Recode2Text(utils.RESPMSG_ERROR_PICTUREADD),
			})
			return
		}

		intAdminId, _ := strconv.Atoi(adminId)
		adminService.SaveAvatarImg(int64(intAdminId), fname)
		context.JSON(iris.Map{
			"status":     utils.RECODE_OK,
			"image_path": fname,
		})

	})

}
