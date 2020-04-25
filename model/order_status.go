package model

//订单状态结构体定义

type OrderStatus struct {
	StatusId   int64  `xorm:"pk autoincr" "json:id"`
	StatusDesc string `xorm:"varchar(255)"`
}
