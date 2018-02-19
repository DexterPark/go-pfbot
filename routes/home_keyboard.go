package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamchanii/go-gin-kakaopf-api-example/pf"
)

// HomeKeyboardRoute 는 최초 시작 시 응답하는 라우트입니다.
func HomeKeyboardRoute(c *gin.Context) {
	c.JSON(http.StatusOK, pf.Keyboard(nil))
}
