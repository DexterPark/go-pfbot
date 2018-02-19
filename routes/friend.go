package routes

import "github.com/gin-gonic/gin"
import "net/http"

// FriendAddedRoute 는 친구 추가 시 사용합니다.
func FriendAddedRoute(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

// FriendRemovedRoute 는 친구 삭제(차단) 시 사용합니다.
func FriendRemovedRoute(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}
