package pfbot

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON 타입은 JSON 방식으로 맵을 활용하기 위해 사용합니다.
type JSON map[string]interface{}

// Request 는 API 호출 시 전달받은 데이터를 사용하기 위한 구조체입니다.
type Request struct {
	UserKey string `json:"user_key" form:"user_key"`
	Type    string `json:"type" form:"type"`
	Content string `json:"content" form:"content"`
}

type routeFunc func(req Request) (interface{}, error)

// Photo 는 Message 맵에서 사용합니다. 이미지에 대한 정보를 담고 있습니다.
type Photo struct {
	url    string
	width  int
	height int
}

// MessageButton 은 Message 맵에서 사용합니다. 메세지 버튼에 대한 정보를 담고 있습니다.
type MessageButton struct {
	label string
	url   string
}

// Keyboard 함수는 Keyboard 구조체 형식의 json 맵을 반환합니다. buttons가 nil 이면 text 형식을 반환하고, []string 인자를 넘기면 buttons 형식으로 반환한다.
func Keyboard(buttons []string) JSON {
	result := JSON{
		"type": "text",
	}

	if buttons != nil {
		result = JSON{
			"type":    "buttons",
			"buttons": buttons,
		}
	}

	return result
}

// Message 함수는 카카오톡 이용자가 명령어를 선택 혹은 입력했을 때 이용자에게 전송하는 응답 메시지 구조체의 json 맵을 반환합니다.
func Message(text string, photo *Photo, messageButton *MessageButton) JSON {
	result := make(JSON)

	if text != "" {
		result["text"] = text
	}

	if photo != nil {
		result["photo"] = JSON{
			"url":    photo.url,
			"width":  photo.width,
			"height": photo.height,
		}
	}

	if messageButton != nil {
		result["message_button"] = JSON{
			"label": messageButton.label,
			"url":   messageButton.url,
		}
	}

	return result
}

// MessageResponse 함수는 OnMessage 에서 반환해야 하는 값을 만들기 위해 사용합니다.
func MessageResponse(message JSON, keyboard JSON) JSON {
	return JSON{
		"message":  message,
		"keyboard": keyboard,
	}
}

// Bot 은 플러스친구 봇 구조체입니다.
type Bot struct {
	onKeyboard routeFunc
	onMessage  routeFunc
	onAdded    routeFunc
	onRemoved  routeFunc
	onLeave    routeFunc
	router     *gin.Engine
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// NewBot 함수는 Bot 구조체를 새로 생성하여 반환합니다.
func NewBot() *Bot {
	bot := new(Bot)
	bot.router = gin.Default()

	return bot
}

func defaultFunc(c *gin.Context) (interface{}, error) {
	return map[string]interface{}{
		"err": c.Request.RequestURI + " is not defined.",
	}, nil
}

func handleFunc(f routeFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var result interface{}
		var err error
		var resp Request

		switch c.Request.Method {
		case "GET":
		case "DELETE":
			resp = Request{
				UserKey: c.Param("user_key"),
			}
		default:
			if err = c.Bind(&resp); err != nil {
				c.AbortWithError(400, err)
				return
			}
		}

		if f != nil {
			result, err = f(resp)
		} else {
			result, err = defaultFunc(c)
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, result)
			return
		}

		c.JSON(http.StatusOK, result)
		return
	}
}

// SetOnKeyboard 는 Bot 구조체의 onKeyboard에 함수를 할당합니다.
func (b *Bot) SetOnKeyboard(f routeFunc) {
	b.onKeyboard = f
	return
}

// SetOnMessage 는 Bot 구조체의 onMessage에 함수를 할당합니다.
func (b *Bot) SetOnMessage(f routeFunc) {
	b.onMessage = f
	return
}

// SetOnAdded 는 Bot 구조체의 onAdded에 함수를 할당합니다.
func (b *Bot) SetOnAdded(f routeFunc) {
	b.onAdded = f
	return
}

// SetOnRemoved 는 Bot 구조체의 onRemoved에 함수를 할당합니다.
func (b *Bot) SetOnRemoved(f routeFunc) {
	b.onRemoved = f
	return
}

// SetOnLeave 는 Bot 구조체의 onLeave에 함수를 할당합니다.
func (b *Bot) SetOnLeave(f routeFunc) {
	b.onLeave = f
	return
}

// Start 함수는 Bot 구조체의 이벤트들을 각 라우터에 할당하고 입력받은 포트 번호로 서버를 실행합니다.
func (b *Bot) Start(port string) {
	// 플러스친구 API에 필요한 라우터를 등록합니다.
	b.router.GET("/keyboard", handleFunc(b.onKeyboard))
	b.router.POST("/message", handleFunc(b.onMessage))
	b.router.POST("/friend", handleFunc(b.onAdded))
	b.router.DELETE("/friend/:user_key", handleFunc(b.onRemoved))
	b.router.DELETE("/chat_room/:user_key", handleFunc(b.onLeave))

	// API 서버를 실행합니다.
	b.router.Run(port)
}
