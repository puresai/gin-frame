package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"fmt"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/13sai/gin-frame/models"
	"github.com/13sai/gin-frame/config"
	"github.com/13sai/gin-frame/router"
	"github.com/13sai/gin-frame/graceful"
	"github.com/13sai/gin-frame/services"
)

var (
	cfg = pflag.StringP("config", "c", "", "config file path.")
)

func main() {
	pflag.Parse()
	// 初始化配置
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// 连接mysql数据库
	issucc := models.GetInstance().InitDataPool()
    if !issucc {
		log.Println("init database pool failure...")
		panic(errors.New("init database pool failure"))
        os.Exit(1)
	}

	// 连接redis
	services.RedisInit()

	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()
	g = router.Load(g)

	go func() {
		if err := pingServer(); err != nil {
			fmt.Println("fail:健康检测失败", err)
		}
		fmt.Println("success:健康检测成功")
	}()

	fmt.Printf("启动http服务端口%s", viper.GetString("addr"))
	fmt.Println()

	if err := graceful.ListenAndServe(viper.GetString("addr"), g); err != nil && err != http.ErrServerClosed {
		fmt.Printf("fail:http服务启动失败: %s\n", err)
	}
}

// 健康检查
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
