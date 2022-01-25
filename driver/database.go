package driver

import (
	news "capstone/backend/features/News/data"
	onlineClass "capstone/backend/features/onlineClass"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	db, err := gorm.Open(mysql.Open("root:@/gym?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = db
	DB.AutoMigrate(&news.News{}, &onlineClass.OnlineClassCore{})
}
