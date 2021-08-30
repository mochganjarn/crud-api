package gorm_client

import (
	"fmt"

	"github.com/rysmaadit/go-template/app"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection(application *app.Application) (db *gorm.DB, er error) {
	s := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		application.Config.DbUser, application.Config.DbPassword,
		application.Config.DbAddr, application.Config.DbPort, application.Config.DbName)

	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
