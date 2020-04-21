package model

import "time"

// 定义管理员结构体
type Admin struct {
	//
	AdminId    int64     `xorm:"pk autoincr" json:"id"`
	AdminName  string    `xorm:"varchar(32)" json:"admin_name"`
	CreateTime time.Time `xorm:"DataTime" json:"create_time"`
	status     int64     `xorm:"default 0" json:"status"`
	Avatar     string    `xorm:"varchar(255)" json:"avatar"`
	Pwd        string    `xorm:"varchar(255)" json:"pwd"`
	CityName   string    `xrom:"varchar(12)" json:"city_name"`
	CityId     int64     `xorm:"index" json:"city_id"`
	Ctiy       *City     `xorm:"- <- ->"` // 所对应的城市结构体
}

// 从Admin数据库实体转换为前端请求的resp的json格式

func (this *Admin) AdminToRespDesc() interface{} {

	respDesc := map[string]interface{}{
		"user_name":   this.AdminName,
		"id":          this.AdminId,
		"create_time": this.CreateTime,
		"status":      this.status,
		"avatar":      this.Avatar,
		"city":        this.CityName,
		"admin":       "管理员",
	}

}
