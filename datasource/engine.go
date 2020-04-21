package datasource

import (
	"github.com/go-xorm/xorm"
	"irisDemo/QianFengCmsProject/model"
)

// 实例化数据库引擎方法：mysql的数据库引擎

func NewMysqlEngine() *xorm.Engine {
	// 创建数据库引擎对象
	engine, err := xorm.NewEngine("MySQL", "root:e13212ss@/Test?charset = utf8")
	if err != nil {
		panic(err.Error())
	}

	//同步数据库结构:主要负责对数据结构实体同步更新到数据库表
	/*
		自动检测和创建表，这个检测是根据表的名字
		自动检测和更新表中的字段，这个检测是根据字段名、同时对表中多余的字段给出警告信息
		自动检测，创建和删除索引和唯一索引，这个检测是根据索引的一个多个字段名字，而不是
		根据索引名称，因此这里需要注意。如果在一个有大量数据的表中引入,,,
		自动转换varchar字段类型到text字段类型。自动警告其他字段类型在模型和数据库之间不一致的情况
		自动警告字段的默认值，是否为空信息在模型和数据库之间不匹配的情况
	*/
	// sync2是sync的基础上优化的方法
	err = engine.Sync2(
		new(model.Permission),
		new(model.City),
		new(model.Admin),
		new(model.AdminPermission),
		new(model.User),
		new(model.UserOrder),
	)
	if err != nil {
		panic(err.Error())
	}
	// 设置显示sql语句
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)
	return engine
}
