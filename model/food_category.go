package model

//定义食品种类结构体
type FoodCategory struct {
	Id               int64  `xorm:"pk autoincr" json:"id"`
	CategoryName     string `json:"name"`
	CategoryDesc     string `json:"description"`
	Level            int64  `json:"level"`
	ParentCategoryId int64  `json:"parent_category_id"`
	RestaurantId     int64  `xorm:"index" json:"restaurant_id"`
}

