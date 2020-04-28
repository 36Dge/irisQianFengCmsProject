package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"irisDemo/QianFengCmsProject/service"
	"irisDemo/QianFengCmsProject/utils"
	"strconv"
)

type OrderController struct {
	Ctx     iris.Context
	Service service.OrderService
	Session *sessions.Session
}

// 获取订单列表
func (orderController *OrderController) Get() mvc.Result {
	iris.New().Logger().Info("查询订单列表")
	offsetStr := orderController.Ctx.FormValue("offset")
	limitStr := orderController.Ctx.FormValue("limit")
	var offset int
	var limit int

	// 判断offset和limit两个变量任意一个都不为""
	if offsetStr == "" || limit == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERLIST,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ORDERLIST),
			},
		}
	}

	offset, err := strconv.Atoi(offsetStr)
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERLIST,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ORDERLIST),
			},
		}
	}

	// 做页数的限制
	if offset <= 0 {
		offset = 0
	}

	//做最大限制
	if limit > MaxLimit {
		limit = MaxLimit
	}

	orderList := orderController.Service.GetOrderList(offset, limit)
	if len(orderList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERLIST,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ORDERLIST),
			},
		}
	}

	// 将查询到的用户数据进行转换成前端需要的内容
	var respList []interface{}
	for _, detail := range orderList {
		respList = append(respList, detail.OrderDetailsResp2())
	}

	// 返回用户列表
	return mvc.Response{
		Object: &respList,
	}

}

// 查询订单记录总数


