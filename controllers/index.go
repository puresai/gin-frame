package controllers

import (
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/13sai/gin-frame/models"
	"github.com/13sai/gin-frame/services"
	"github.com/13sai/gin-frame/logging"
) 

func AddIntegral(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	logging.Info((c.Request.Body))
	logging.Info((data))

	target := models.InteagralInputJson{}
	if err := json.Unmarshal([]byte(data), &target); err != nil {
		SendResponse(c, 1001, "数据格式或类型错误", nil)
		logging.Error(target)
		return
	}

	if target.IntegralCount < 1 {
		SendResponse(c, 1003, "积分不能小于0", nil)
		logging.Error(target)
		return
	}

	if target.AccountId < 1 {
		SendResponse(c, 1003, "用户id参数错误", nil)
		logging.Error(target)
		return
	}

	str := strings.ToLower(fmt.Sprintf("%s%d%s%d", "account_id", target.AccountId, "count", target.IntegralCount))
	if services.HmacSha1(str) != target.Key {
		SendResponse(c, 1002, "数据格式错误", nil)
		logging.Error(target)
		return
	}

	score, err := services.AddIntegral(target)

	if err != nil {
		SendResponse(c, 2001, "积分增加失败", nil)
		logging.Info(target)
		logging.Error(err)
		return
	}

	// 积分写入redis
	go func() {
		target.AddTime = time.Now().Unix()
		jsonS, err := json.Marshal(target)
		if err != nil {
			logging.Info(target)
			logging.Error(err)
			return
		}
		services.LPush(services.IntegralRecord, jsonS)
	}()

	SendResponse(c, 0, "success", models.IntegralData{Integral:int(score)})
	return
}

func GetIntegral(c *gin.Context) {
	accountId := c.Param("accountId")
	integral := services.GetIntegral(accountId)
	SendResponse(c, 0, "success", models.IntegralData{Integral:integral})
}

func Health(c *gin.Context) {
	SendResponse(c, 0, "success", nil)
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}