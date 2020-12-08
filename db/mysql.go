package db

import (
    "fmt"
    "sync"
    "errors"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/spf13/viper"
)

var once sync.Once
var err error 
var db *gorm.DB

// 单例模式
func GetDB() *gorm.DB {
    once.Do(func() {
        db = openPool()
    })

    return db
}

func openPool() *gorm.DB {
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.host"), viper.GetString("db.name"), viper.GetString("db.charset"))
    db, err := gorm.Open("mysql", dsn)
    if err != nil {
        panic(errors.New("mysql连接失败"))
    }

    db.DB().SetMaxIdleConns(viper.GetInt("db.MaxIdleConns"))
    db.DB().SetMaxOpenConns(viper.GetInt("db.MaxOpenConns"))
	if viper.GetBool("db.log") {
		db.LogMode(true)
	}
    return db
}

func CloseDB(db *gorm.DB)  {
	db.Close()
}