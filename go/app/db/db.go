package db

import (
	"fmt"

	"../model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

var err error

func Connect() {

	CONNECT := "root:@tcp(127.0.0.1:3306)/test_chat?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	// データベースに接続する
	db, err = gorm.Open("mysql", CONNECT)
	// データベースに接続できなかった場合の処理
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("データベースに接続しました")
	}

	// db.DropTable(&model.User{}, &model.UserMessage{}, &model.Userchatroom{})
	db.AutoMigrate(&model.User{}, &model.UserMessage{}, &model.Userchatroom{})

	db.LogMode(true)

}

// データベースの情報を取得する
func Get() *gorm.DB {
	return db
}

// データベースの接続を切断する
func Close() {
	db.Close()
}
