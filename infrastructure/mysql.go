package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitMySQL() (err error) {
	Db, err = gorm.Open(mysql.New(mysql.Config{
		//DSN:        "admin:1234@tcp(mysql80:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // data source name, 详情参考：https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DSN: "admin:1234@tcp(localhost:3306)/be103?charset=utf8mb4&parseTime=True&loc=Local", // data source name, 详情参考：https://github.com/go-sql-driver/mysql#dsn-data-source-name
	}), &gorm.Config{})
	return
}
