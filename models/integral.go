package models

type IntegralData struct {
    Integral int `json:"integral"`
}

type InteagralInputJson struct {
    AccountId int64 `json:"account_id"`
	IntegralCount int `json:"count"`
	Key string `json:"key"`
	AddTime int64 `json:"add_time"`
}
