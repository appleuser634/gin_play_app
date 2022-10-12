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
	UserAuth(authInfo) bool
	SendMessage(requests.SendMessageRequest) int
	GetMessage(requests.GetMessageRequest) int
}

type authInfo struct {
	name  string
	token string
}

func NewController(model model.Model) Controller {
	return &controller{model}
}

func (controller *controller) GetUserName() {

	userlist := controller.model.GetUserName()

	fmt.Println("Called Controller!!")
	fmt.Println(userlist[0].Name)
}

func (controller *controller) UserAuth(auth authInfo) bool {

	// ユーザのTokenを取得
	userlist := controller.model.GetUserToken(auth.name)

	// 該当ユーザーがいなければfalseを返す
	if len(userlist) == 0 {
		fmt.Println("Auth Fail!!")
		return false
	}

	fmt.Printf("User Token:%v\n", userlist[0].Token)

	// Tokenを照合して認証
	if auth.token == userlist[0].Token {
		fmt.Println("Auth OK!!!")
		return true
	} else {
		fmt.Println("Auth FAIL!!!")
	}

	return false
}

func (controller *controller) SendMessage(SendMessageJson requests.SendMessageRequest) int {

	auth := authInfo{}

	auth.name = SendMessageJson.From
	auth.token = SendMessageJson.Token

	authResult := controller.UserAuth(auth)

	// 認証失敗であれば400エラーを返す
	if !authResult {
		return http.StatusBadRequest
	}

	// 送信メッセージを登録する
	controller.model.CreateMessage(SendMessageJson.Message, SendMessageJson.From, SendMessageJson.To)

	return http.StatusOK
}

func (controller *controller) GetMessage(GetMessageJson requests.GetMessageRequest) int {

	auth := authInfo{}

	auth.name = GetMessageJson.To
	auth.token = GetMessageJson.Token

	authResult := controller.UserAuth(auth)

	// 認証失敗であれば400エラーを返す
	if !authResult {
		return http.StatusBadRequest
	}

	// 受信メッセージを取得する
	messagelist := controller.model.GetMessage(GetMessageJson.MessageID, GetMessageJson.To)

	for _, m := range messagelist {
		fmt.Printf("MessageID:%d\n", m.ID)
		fmt.Printf("Message From:%s\n", m.MessageFrom)
		fmt.Printf("Message:%s\n", m.Message)
	}

	return http.StatusOK

}
