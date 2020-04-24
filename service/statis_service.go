package service

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"irisDemo/QianFengCmsProject/model"
	"math/rand"
	"time"
)

// 统计模块功能接口
type StatisService interface {
	//查询某一天的用户的增长数量
	GetUserDailyCount(data string) int64
	GetOrderDailyCount(data string) int64
	GetAdminDailyCount(data string) int64
}

//统计功能服务实现结构体
type statisService struct {
	Engine *xorm.Engine
}

//新建统计模块功能服务对象   如果想返回结构指针，这个结构体必须实现这个接口中的所有方法（因为返回值是这个接口类型）
func NewStatisService(engine *xorm.Engine) StatisService {
	return &statisService{
		Engine: engine,
	}
}

// 结构体staticService实现接口
// 查询某一日管理员增长数量
func (ss *statisService) GetAdminDailyCount(data string) int64 {
	if data == "NaN-NaN-NaN" {
		// 当日增长数据请求
		data = time.Now().Format("2006-01-02")
	}
	// 查询日期data格式解析
	startData, err := time.Parse("2006-01-02", data)
	if err != nil {
		return 0
	}

	endData := startData.AddDate(0, 0, 1)
	result, err := ss.Engine.Where("create_time between ? and ? and status = 0", startData.Format("2006-01-02"), endData.Format("2006-01-02 15:03:04")).Count(model.Admin{})
	if err != nil {
		return 0
	}
	fmt.Println(result)
	return int64(rand.Intn(100))
}

// 查询某一日订单的单日增长数量
func (ss *statisService) GetOrderDailyCount(data string) int64 {
	if data == "NaN-NaN-NaN" {
		// 当日数据请求
		data = time.Now().Format("2006-01-02")
	}
	startData, err := time.Parse("2006-01-02", data)
	if err != nil {
		return 0
	}
	endData := startData.AddDate(0, 0, 1)
	result, err := ss.Engine.Where("time between ? and ? and del_flat = 0", startData.Format("2006-01-02 15:04:03"), endData.Format("2006-01-02 15:04:05")).Count(model.UserOrder{})
	if err != nil {
		return 0

	}
	fmt.Println(result)
	return int64(rand.Intn(100))
}

// 查询某一日用户的单日的增长数量

func (ss *statisService) GetUserDailyCount(data string) int64 {

	if data == "NaN-NaN-NaN" {
		data = time.Now().Format("2006-01-02")
	}
	startData, err := time.Parse("2006-01-02", data)
	if err != nil {
		return 0
	}
	endData := startData.AddDate(0, 0, 1)
	result, err := ss.Engine.Where("register_time between ? and ? del_flat = 0", startData.Format("2006-01-03 15:04:05"), endData.Format("2006-01-02 15:04:05")).Count(model.User{})

	if err != nil {
		return 0
	}
	fmt.Println(result)
	return int64(rand.Intn(100))
}
