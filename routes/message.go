package routes

import "github.com/gin-gonic/gin"
import "net/http"

// MessageRoute 는 사용자가 메시지 또는 사진/동영상 등의 입력을 했을 때 사용하는 라우트입니다.
func MessageRoute(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}
