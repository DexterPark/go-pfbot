# go-pfbot
카카오 플러스친구 자동응답 API 봇 라이브러리 based on [Gin](https://github.com/gin-gonic/gin)

# Installation
```
go get github.com/iamchanii/go-pfbot
```

# Quick Start
```go
package main

import "github.com/iamchanii/go-pfbot"

func main() {
  bot := pfbot.NewBot()

  bot.SetOnKeyboard(func(req pfbot.Request) (interface{}, error) {
    return pfbot.Keyboard([]string{"시작하기"}), nil
  })

  bot.Start(":8080")
}
```
```bash
$ curl -X GET http://localhost:8080/keyboard
```

# Example
- [go-pfbot-example](https://github.com/iamchanii/go-pfbot-example)

# License
- MIT License