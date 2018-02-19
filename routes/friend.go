package routes

import "github.com/gin-gonic/gin"
import "net/http"

func FriendAddedRoute(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func FriendRemovedRoute(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}
