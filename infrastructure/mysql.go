package infrastructure

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitMySQL() (err error) {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC", config.AppConfig.AwsDbUsername, config.AppConfig.AwsDbPassword, config.AppConfig.AwsDbHost, "3306", "shopping_cart")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC", "admin", "1234", "localhost", "3306", "be103") // 這還好我就寫死沒差
	Db, err = gorm.Open(mysql.New(mysql.Config{
		//DSN:        "admin:1234@tcp(mysql80:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // data source name, 详情参考：https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DSN: dsn, // data source name, 详情参考：https://github.com/go-sql-driver/mysql#dsn-data-source-name
	}), &gorm.Config{})
	return
}
