package model

//用户订单详情结构体
type OrderDetail struct {
	UserOrder   `xrom:"extends"`
	User        `xrom:"extends"`
	OrderStatus `xrom:"extends"`
	Shop        `xrom:"extends"`
	Address     `xrom:"extends"`
}

func (detail *OrderDetail) OrderDetailsResp2() interface{} {
	respDesc := map[string]interface{}{
		"id":                    detail.UserOrder.Id,
		"total_amount":          detail.UserOrder.SumMoney,
		"user_id":               detail.User.UserName,
		"status":                detail.OrderStatus.StatusDesc,
		"restaurant_id":         detail.Shop.ShopId,
		"restaurant_image_url":  detail.Shop.ImagePath,
		"restaurant_created_at": detail.Time,
		"status_code":           0,
		"address_id":            detail.Address.AddressId,
	}

	statusDesc := map[string]interface{}{
		"color":     "f60",
		"sub_title": "15分钟提交",
		"title":     detail.OrderStatus.StatusDesc,
	}
	respDesc["status_bar"] = statusDesc
	return respDesc

}
