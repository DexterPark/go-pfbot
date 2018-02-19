package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iamchanii/go-gin-kakaopf-api-example/routes"
)

func main() {
	// 기본 미들웨어를 적용한 gin 라우터를 생성합니다.
	router := gin.Default()

	// 플러스친구 API에 필요한 라우터를 등록합니다.
	router.GET("/keyboard", routes.HomeKeyboardRoute)
	router.POST("/message", routes.MessageRoute)
	router.POST("/friend", routes.FriendAddedRoute)
	router.DELETE("/friend", routes.FriendRemovedRoute)
	router.DELETE("/chat_room/:user_key", routes.LeaveChatRoomRoute)

	// API 서버를 실행합니다.
	router.Run()
}
