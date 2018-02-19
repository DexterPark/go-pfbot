package routes

import "github.com/gin-gonic/gin"
import "net/http"

func MessageRoute(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}
