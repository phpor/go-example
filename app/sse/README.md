# SSE 示例项目

## 功能说明
- 服务端每2秒发送时间戳事件
- 客户端实时显示接收到的SSE事件
- 通过Go提供静态文件服务

## 实现清单
- [x] 需求确认：基础SSE实现（时间戳事件+文本展示）
- [x] 创建服务端代码：app/sse/main.go
- [x] 创建静态资源目录：app/sse/static/
- [x] 编写HTML页面：app/sse/static/index.html  
- [x] 实现SSE服务端逻辑（Go HTTP服务+事件流）
- [x] 实现SSE客户端逻辑（JS EventSource）
- [x] 添加静态文件服务路由（Go http.FileServer）
- [x] 测试端到端功能

## 使用说明
1. 启动服务：`go run main.go`
2. 访问客户端：http://localhost:8080