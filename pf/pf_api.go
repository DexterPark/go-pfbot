package pf

// JSON 은 gin.H 와 같은 형식의 타입을 선언합니다.
type JSON map[string]interface{}

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
