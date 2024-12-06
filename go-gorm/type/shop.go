package _type

import (
	"gorm.io/gorm"
	"time"
)

type TbShop struct {
	ID         uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string     `gorm:"type:varchar(128);not null" json:"name"`
	TypeID     uint64     `gorm:"not null" json:"type_id"`
	Images     string     `gorm:"type:varchar(1024);not null" json:"images"`
	Area       string     `gorm:"type:varchar(128)" json:"area"`
	Address    string     `gorm:"type:varchar(255);not null" json:"address"`
	X          float64    `gorm:"not null" json:"x"`
	Y          float64    `gorm:"not null" json:"y"`
	AvgPrice   *uint64    `gorm:"" json:"avg_price"`
	Sold       string     `gorm:"type:int(10) unsigned zerofill;not null" json:"sold"`
	Comments   string     `gorm:"type:int(10) unsigned zerofill;not null" json:"comments"`
	Score      string     `gorm:"type:int(2) unsigned zerofill;not null" json:"score"`
	OpenHours  string     `gorm:"type:varchar(32)" json:"open_hours"`
	CreateTime *time.Time `gorm:"type:timestamp" json:"create_time"`
	UpdateTime *time.Time `gorm:"type:timestamp" json:"update_time"`
}

func NewShop() *TbShop {
	return &TbShop{}
}

func (t TbShop) Create(db *gorm.DB) {
	newShop := TbShop{
		Name:      "示例商铺",
		TypeID:    1,
		Images:    "image1.jpg,image2.jpg",
		Address:   "示例地址",
		X:         120.0,
		Y:         30.0,
		Sold:      "0000000000",
		Comments:  "0000000000",
		Score:     "05",
		OpenHours: "09:00-18:00",
	}
	res := db.Create(&newShop).Table("tb_shop")
	if res.Error != nil {
		panic(res.Error)
	}
}

func (t TbShop) Find(db *gorm.DB) {
	newShop := TbShop{
		Name:      "示例商铺",
		TypeID:    1,
		Images:    "image1.jpg,image2.jpg",
		Address:   "示例地址",
		X:         120.0,
		Y:         30.0,
		Sold:      "0000000000",
		Comments:  "0000000000",
		Score:     "05",
		OpenHours: "09:00-18:00",
	}
	// 根据主键查询第一条记录
	db.First(&newShop).Table("tb_shop")
	//// SELECT * FROM users ORDER BY id LIMIT 1;

	// 随机获取一条记录
	db.Take(&newShop).Table("tb_shop")
	//// SELECT * FROM users LIMIT 1;

	// 根据主键查询最后一条记录
	db.Last(&newShop).Table("tb_shop")
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	// 查询所有的记录
	db.Find(&newShop).Table("tb_shop")
	//// SELECT * FROM users;
}
