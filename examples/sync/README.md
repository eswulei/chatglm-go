# Sync

To run an example:

```shell
export API_KEY="<your api key here>"
go run ./example/async/
```

```go
package main

import (
	"fmt"
	"github.com/eswulei/chatglm-go"
	"os"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	m := chatglm.ModelAPI{
		Model:       chatglm.ChatGLMLite,
		TopP:        0.7,
		Temperature: 0.9,
		Prompt: []map[string]interface{}{
			{"role": "user", "content": "你好"},
			{"role": "assistant", "content": "我是人工智能助手"},
			{"role": "user", "content": "你叫什么名字"},
			{"role": "assistant", "content": "我叫chatGLM"},
			{"role": "user", "content": "你都可以做些什么事"},
		},
	}
	resp, err := m.Invoke(apiKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
```