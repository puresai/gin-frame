package services

import (
	"fmt"
    "crypto/hmac"
	"crypto/sha1"
	"strconv"
	// "errors"

	"github.com/13sai/gin-frame/models"
	// "github.com/13sai/gin-frame/logging"
)

const (
	AccountIdKey string = "list"
	UserIdAccount string = "user_id"
	IntegralRecord string = "record"
	Salt string = "y4#14%sf@p";
)

//通过accountId获取积分
func GetIntegral(accountId string) (integral int ){
	score := RedisClient.ZScore(AccountIdKey, accountId)
	if score.Err() != nil {
		integral = 0
		return
	} else {
		ret := score.Val()
		integral = int(ret)
		return 
	}
}

//通过accountId添加积分
// {
//     "account_id":27117671, 
//     "count":10,
//     "key": "c665e55287485bd6eb586bf7ac06c3be418f7ede"
// }
func AddIntegral(data models.InteagralInputJson) (score float64, err error) {
	score, err = RedisClient.ZIncrBy(AccountIdKey, float64(data.IntegralCount), strconv.FormatInt(data.AccountId, 10)).Result()
	return 
}

func HmacSha1(str string) string {
	key := []byte(Salt)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(str))
	return fmt.Sprintf("%x", mac.Sum(nil))
}