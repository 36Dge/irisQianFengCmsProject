package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"go/ast"
	"strconv"
)

//食品类型控制器
type CategoryController struct {
	Ctx     iris.Context
	Service service.CategoryService
}
type CategoryEntity struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	RestaruantId string `json:"restaruant_id"`
}

func (cc *CategoryController) BeforeActivation(a mvc.BeforeActivation) {

	// 通过商铺Id获取对应的食品种类
	a.Handle("GET", "/gatcategory/{shopId}", "GetCategoryByShopId")
	// 获取全部的食品种类
	a.Handle("GET", "/v2/restaruant/category", "GetAllCategory")
	//添加商铺记录
	a.Handle("POST", "/addShop", "PostAddShop")
	a.Handle("DELETE", "/restaurant/{restaurant_id}", "DeleteRestaurant")

}

//删除商户记录
func (cc *CategoryController) DeleteRestaurant() mvc.Result {
	restaurant_id := cc.Ctx.Params().Get("restaurant_id")
	shopId,err := strconv.Atoi(restaurant_id)
	if err != nil{
		iris.New().Logger().Info(err.Error())

	}

}
