package models

import (
	"io/ioutil"
	"net/http"

	"github.com/widuu/gojson"
)

type Users struct {
	ID       int    `gorm:"column:id;primary_key"`
	Name     string `gorm:"column:name"`
	Password string `gorm:"column:password"`
	Token    string `gorm:"column:token"`
}

func (user *Users) GetToken() (token string, err error) {
	key := "dingzoaqctxwilo7na04"
	secret := "05Xio0tV8rn29Q0xCoyLqhf0dYXZsjmdnQSotvfwAyLh89oFYuvrj6umRzK_rsCO"
	tokenUrl := "https://oapi.dingtalk.com/gettoken?appkey=" + key + "&appsecret=" + secret
	client := &http.Client{}
	resp, err := client.Get(tokenUrl)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	accessToken := gojson.Json(string(body)).Get("access_token").Tostring()
	return accessToken, nil
}

func (user *Users) GetUserId(code string, accessToken string) (bodys []byte, err error) {
	userInfoUrl := "https://oapi.dingtalk.com/user/getuserinfo?access_token=" + accessToken + "&code=" + code
	client := &http.Client{}
	resp, err := client.Get(userInfoUrl)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return bodys, err
	}
	return body, nil
}
