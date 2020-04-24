package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"irisDemo/QianFengCmsProject/service"
	"irisDemo/QianFengCmsProject/utils"
	"strings"
)

// 统计功能控制者

type StatisController struct {
	// 上下文对象
	Ctx iris.Context

	//统计功能的服务实现接口
	Service service.StatisService

	//session
	Session *sessions.Session
}

var (
	ADMINMODULE = "ADMIN_"
	USERMODULE  = "USER_"
	ORDERMODULE = "ORDER_"
)

// 解析统计功能路由请求

func (sc *StatisController) GetCount() mvc.Result {
	path := sc.Ctx.Path()
	var pathSlice []string
	if path != "" {
		// pathSlice = {"" "static" "user" "2019-03-10" "count"}
		pathSlice = strings.Split(path, "/")
	}
	//不符合请求格式
	if len(pathSlice) != 5 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_FAIL,
				"count":  0,
			},
		}
	}

	// 将第一个元素去掉
	pathSlice = pathSlice[1:]
	model := pathSlice[1]
	data := pathSlice[2]
	var result int64
	switch model {
	case "user":
		userResult := sc.Session.Get(USERMODULE + data)
		if userResult != nil {
			userResult = userResult.(float64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": utils.RECODE_OK,
					"count":  userResult,
				},
			}
		} else {
			iris.New().Logger().Error(data) // 时间
			result = sc.Service.GetUserDailyCount(data)
			//设置缓存
			sc.Session.Set(USERMODULE+data, result)
		}

	case "order":
		orderStatis := sc.Session.Get(ORDERMODULE + data)
		if orderStatis != nil {
			orderStatis = orderStatis.(float64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": utils.RECODE_OK,
					"count":  orderStatis,
				},
			}
		} else {
			result = sc.Service.GetOrderDailyCount(data)
			sc.Session.Set(ORDERMODULE+data, result)
		}
	case "admin":
		adminStatis := sc.Session.Get(ADMINMODULE + data)
		if adminStatis != nil {
			adminStatis = adminStatis.(float64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": utils.RECODE_OK,
					"count":  adminStatis,
				},
			}
		} else {
			result = sc.Service.GetAdminDailyCount(data)
			sc.Session.Set(ADMINMODULE, result)
		}

	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count":  result,
		},
	}

}
