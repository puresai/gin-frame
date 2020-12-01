package main

import (
	"fmt"
	"time"
	"strconv"
	"net/http"

    "github.com/spf13/pflag"
    "github.com/spf13/viper"
    "github.com/gin-gonic/gin"

    "github.com/13sai/gin-frame/config"
    "github.com/13sai/gin-frame/db"
    "github.com/13sai/gin-frame/router"
    "github.com/13sai/gin-frame/logger"
    "github.com/13sai/gin-frame/graceful"
)

var (
    conf = pflag.StringP("config", "c", "", "config filepath")
)

func main() {
    pflag.Parse()

    // 初始化配置
    if err := config.Run(*conf); err != nil {
        panic(err)
	}

	logger.Info("i'm log123-----Info")
	logger.Error("i'm log123-----Error")

	
	// 连接mysql数据库
	DB := db.GetDB()
	defer db.CloseDB(DB)

	// redis
	db.RedisInit()

	go func() {
		pingServer()
	}()

	gin.SetMode(viper.GetString("mode"))
	g := gin.New()
	g = router.Load(g)

	// g.Run(viper.GetString("addr"))


	// logger.Info("启动http服务端口%s\n", viper.GetString("addr"))

	// time.Sleep(2*time.Second)
	if err := graceful.ListenAndServe(viper.GetString("addr"), g); err != nil && err != http.ErrServerClosed {
		logger.Error("fail:http服务启动失败: %s\n", err)
	}
}

// 健康检查
// func pingServer() error {
// 	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
// 		url := fmt.Sprintf("%s%s%s", "http://127.0.0.1", viper.GetString("addr"), viper.GetString("healthCheck"))
// 		fmt.Println(url)
// 		resp, err := http.Get(url)
// 		if err == nil && resp.StatusCode == 200 {
// 			return nil
// 		}
// 		time.Sleep(time.Second)
// 	}
// 	return errors.New("健康检测404")
// }

// 健康检查
func pingServer() {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		url := fmt.Sprintf("%s%s%s", "http://127.0.0.1", viper.GetString("addr"), viper.GetString("healthCheck"))
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			fmt.Println("health check success!")
			return
		}
		fmt.Println("check fail -" + strconv.Itoa(i+1)+"times")
		time.Sleep(time.Second)
	}
	fmt.Println("Cannot connect to the router!!!")
}