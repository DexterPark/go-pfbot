package routes

import "github.com/gin-gonic/gin"
import "net/http"

// LeaveChatRoomRoute 는 사용자가 채팅방을 목록에서 삭제했을 경우 사용됩니다.
func LeaveChatRoomRoute(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}
