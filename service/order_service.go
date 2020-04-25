package service

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"irisDemo/QianFengCmsProject/model"
)

// 订单服务接口
type OrderService interface {
	GetCount() (int64, error)
	GetOrderList(offset, limit int) []model.OrderDetail
}

// 订单服务
type orderService struct {
	Engine *xorm.Engine
}

// 实例化OrderService服务对象
// 这是一个函数
func NewOrderServic(db *xorm.Engine) OrderService {
	return &orderService{Engine: db} // orderService 必须要实现 接口中方法

}

// 获取订单列表
func (orderService *orderService) GetOrderList(offset, limit int) []model.OrderDetail {
	orderList := make([]model.OrderDetail, 0)

	// 查询用户订单信息 多表关联查询
	err := orderService.Engine.Table("user_order").
		Join("INNER", "order_status", "order_status.status_id = user_order.order_status_id").
		Join("INNER", "user", "user.id = user_order.user_id").
		Join("INNER", "shop", "shop.shop_id = user_order.shop_id").
		Join("INNER", "adderss", "address.address_id = user_order.address_id").
		Find(&orderList)
	iris.New().Logger().Info(orderList[0])
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return orderList
}

