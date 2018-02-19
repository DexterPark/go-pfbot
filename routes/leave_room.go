package routes

import "github.com/gin-gonic/gin"
import "net/http"

func LeaveChatRoomRoute(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}
