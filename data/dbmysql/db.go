package dbmysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "wecom:ZHPtC2EMGa6NeeP8@tcp(124.221.95.39:3306)/wecom?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func likeWrap(content string) string {
	return "%" + content + "%"
}

type Page struct {
	PageNumber int
	PageSize   int
}
