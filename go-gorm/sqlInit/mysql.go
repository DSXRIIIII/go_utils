package sqlInit

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func MySQLDBInit() (*gorm.DB, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/hmdp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})

	if err != nil {
		return nil, err
	}
	return db, nil
}
