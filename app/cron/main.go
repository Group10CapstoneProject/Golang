package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Group10CapstoneProject/Golang/config"
	"github.com/Group10CapstoneProject/Golang/database"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("test")
	config.InitConfig()

	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	t := time.Now()
	scheduler := cron.New(cron.WithLocation(t.Location()))
	fmt.Println("test")
	// stop scheduler tepat sebelum fungsi berakhir
	defer scheduler.Stop()

	// set task yang akan dijalankan scheduler
	// gunakan crontab string untuk mengatur jadwal
	scheduler.AddFunc("* 00 * * *", func() { UpdateMembers(db) })
	scheduler.AddFunc("* 00 * * *", func() { UpdateOfflineClassBooking(db) })
	scheduler.AddFunc("* 00 * * *", func() { UpdateOnlineClassBooking(db) })
	scheduler.AddFunc("* 00 * * *", func() { UpdateTrainerBooking(db) })

	// start scheduler
	go scheduler.Start()

	// trap SIGINT untuk trigger shutdown.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

func UpdateMembers(db *gorm.DB) {
	err := db.Table("members").
		Where("DATE(expired_at) < ? AND status IN ?", time.Now().Format("2006-01-02"), []string{"WAITING", "PENDING"}).
		Update("status", "INACTIVE").
		Error
	if err != nil {
		fmt.Println(err)
	}
	err = db.Table("members").
		Where("DATE(expired_at) < ? AND status = ?", time.Now().Format("2006-01-02"), "ACTIVE").
		Update("status", "DONE").
		Error
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateOfflineClassBooking(db *gorm.DB) {
	err := db.Table("offline_class_bookings").
		Where("DATE(expired_at) < ? AND status IN ?", time.Now().Format("2006-01-02"), []string{"WAITING", "PENDING"}).
		Update("status", "INACTIVE").
		Error
	if err != nil {
		fmt.Println(err)
	}
	err = db.Table("offline_class_bookings").
		Where("DATE(expired_at) < ? AND status = ?", time.Now().Format("2006-01-02"), "ACTIVE").
		Update("status", "DONE").
		Error
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateOnlineClassBooking(db *gorm.DB) {
	err := db.Table("online_class_bookings").
		Where("DATE(expired_at) < ? AND status IN ?", time.Now().Format("2006-01-02"), []string{"WAITING", "PENDING"}).
		Update("status", "INACTIVE").
		Error
	if err != nil {
		fmt.Println(err)
	}
	err = db.Table("online_class_bookings").
		Where("DATE(expired_at) < ? AND status = ?", time.Now().Format("2006-01-02"), "ACTIVE").
		Update("status", "DONE").
		Error
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateTrainerBooking(db *gorm.DB) {
	err := db.Table("trainer_bookings").
		Where("DATE(expired_at) < ? AND status IN ?", time.Now().Format("2006-01-02"), []string{"WAITING", "PENDING"}).
		Update("status", "INACTIVE").
		Error
	if err != nil {
		fmt.Println(err)
	}
	err = db.Table("trainer_bookings").
		Where("DATE(expired_at) < ? AND status = ?", time.Now().Format("2006-01-02"), "ACTIVE").
		Update("status", "DONE").
		Error
	if err != nil {
		fmt.Println(err)
	}
}
