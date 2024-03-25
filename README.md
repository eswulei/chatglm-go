> **Synchronization with [official Python SDK](https://pypi.org/project/zhipuai/) feature & version numbers.**

This library provides unofficial Go clients for [ChatGLM API](https://open.bigmodel.cn/dev/api). We support:

- [Get Token](https://open.bigmodel.cn/usercenter/apikeys)
- [ChatGLM](https://open.bigmodel.cn)

## Installation:

```
go get github.com/eswulei/chatglm-go
```

## Select Model

- ChatGLMLite
- ChatGLMStd
- ChatGLMPro

```go
m := chatglm.ModelAPI{
    Model:       chatglm.ChatGLMLite,
    // Model:       chatglm.ChatGLMStd,
    // Model:       chatglm.ChatGLMPro,
}
```

## Examples

- [sync](https://github.com/eswulei/chatglm-go/tree/main/examples/sync)
- [async](https://github.com/eswulei/chatglm-go/tree/main/examples/async)
- [stream](https://github.com/eswulei/chatglm-go/tree/main/examples/stream)

## license

[Apache License 2.0](./LICENSE)
