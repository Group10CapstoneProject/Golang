package database

import (
	"fmt"

	"github.com/Group10CapstoneProject/Golang/config"
	"github.com/Group10CapstoneProject/Golang/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Env.DB_USERNAME,
		config.Env.DB_PASSWORD,
		config.Env.DB_ADDRESS,
		config.Env.DB_NAME,
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
		model.PaymentMethod{},
		model.MemberType{},
		model.Notification{},
		model.Member{},
		model.OnlineClassCategory{},
		model.OnlineClass{},
		model.OnlineClassBooking{},
		model.OfflineClassCategory{},
		model.OfflineClass{},
		model.OfflineClassBooking{},
	)
}
