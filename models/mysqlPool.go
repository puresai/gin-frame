package models

import (
    "fmt"
	"sync"
	"errors"
	
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

/*
* MysqlPool
* 数据库连接操作库
* 基于gorm封装开发
 */
type MysqlPool struct {}

var instance *MysqlPool
var once sync.Once

var db *gorm.DB
var err_db error

func GetInstance() *MysqlPool {
    once.Do(func() {
        instance = &MysqlPool{}
    })
    return instance
}

/*
* @fuc 初始化数据库连接
*/
func (m *MysqlPool) InitDataPool() (issucc bool) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.host"), viper.GetString("db.name"), viper.GetString("db.charset"))
    db, err_db = gorm.Open("mysql", dsn)
    if err_db != nil {
		panic(errors.New("mysql连接失败"))
        fmt.Println(err_db)
        return false
	}
	db.DB().SetMaxIdleConns(100)
    db.DB().SetMaxOpenConns(200)
    // db.LogMode(true)
    return true
}

/*
* @fuc  对外获取数据库连接对象db
*/
func (m *MysqlPool) GetMysqlDB() (db_con *gorm.DB) {
    return db
}