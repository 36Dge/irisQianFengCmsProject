package model

// 地址结构体实体
type Address struct {
	AddressId     int64  `xorm:"pk autoincr json:"id""`
	Address       string `json:"address"`
	Phone         string `json:"phone"`
	AddressDetail string `json:"address_detail"`
	IsValid       int    `json:"is_valid"`
}
