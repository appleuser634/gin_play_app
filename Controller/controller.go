package controller

import (
	"fmt"
	"net/http"

	"local.package/model"
	"local.package/requests"
)

type controller struct {
	model model.Model
}

type Controller interface {
	GetUserName()
	UserAuth(requests.SendMessageRequest) bool
	SendMessage(requests.SendMessageRequest) int
}

func NewController(model model.Model) Controller {
	return &controller{model}
}

func (controller *controller) GetUserName() {

	userlist := controller.model.GetUserName()

	fmt.Println("Called Controller!!")
	fmt.Println(userlist[0].Name)
}

func (controller *controller) UserAuth(SendMessageJson requests.SendMessageRequest) bool {

	// ユーザのTokenを取得
	userlist := controller.model.GetUserToken(SendMessageJson.From)

	// 該当ユーザーがいなければfalseを返す
	if len(userlist) == 0 {
		fmt.Println("Auth Fail!!")
		return false
	}

	fmt.Printf("User Token:%v\n", userlist[0].Token)

	// Tokenを照合して認証
	if SendMessageJson.Token == userlist[0].Token {
		fmt.Println("Auth OK!!!")
		return true
	} else {
		fmt.Println("Auth FAIL!!!")
	}

	return false
}

func (controller *controller) SendMessage(SendMessageJson requests.SendMessageRequest) int {

	authResult := controller.UserAuth(SendMessageJson)

	// 認証失敗であれば400エラーを返す
	if !authResult {
		return http.StatusBadRequest
	}

	// 送信メッセージを登録する
	controller.model.CreateMessage(SendMessageJson.Message, SendMessageJson.From, SendMessageJson.To)

	return http.StatusOK

	// username := userlist[0].Name
}
