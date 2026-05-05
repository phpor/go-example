package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/v1/chat/completions", handleChatCompletions)

	fmt.Println("✅ OpenAI ↔ Anthropic 代理启动成功！")
	fmt.Println("👉 监听地址：http://localhost:16888")
	fmt.Println("👉 Zed 里配置 api_url = http://localhost:16888/v1")
	if err := http.ListenAndServe(":16888", nil); err != nil {
		fmt.Fprintf(os.Stderr, "❌ 服务启动失败: %v\n", err)
		os.Exit(1)
	}
}

// 接收 Zed 发来的 OpenAI 格式 → 转发成 Anthropic 格式 → 返回 OpenAI 格式
func handleChatCompletions(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		http.Error(w, "missing authorization", 401)
		return
	}
	apiKey := strings.TrimPrefix(auth, "Bearer ")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body: "+err.Error(), 500)
		return
	}

	var openAIReq OpenAIRequest
	if err := json.Unmarshal(body, &openAIReq); err != nil {
		msg := fmt.Sprintf("failed to parse request: %v, body: %s", err, string(body))
		fmt.Println("❌", msg)
		http.Error(w, msg, 400)
		return
	}
	fmt.Printf("📥 收到请求: model=%s, messages=%d\n", openAIReq.Model, len(openAIReq.Messages))

	// 1. 转成 Anthropic 请求
	anthropicReq := convertOpenAIToAnthropic(&openAIReq)
	anthropicReqData, _ := json.Marshal(anthropicReq)

	// 2. 转发到你的 Anthropic 兼容接口
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://127.0.0.1:3456/v1/messages", bytes.NewBuffer(anthropicReqData))
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	// 3. 处理流式或非流式响应
	if openAIReq.Stream {
		handleStreamResponse(w, resp)
		return
	}

	// 非流式响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "failed to read anthropic response: "+err.Error(), 500)
		return
	}

	var anthropicResp AnthropicResponse
	if err := json.Unmarshal(respBody, &anthropicResp); err != nil {
		msg := fmt.Sprintf("failed to parse anthropic response: %v, body: %s", err, string(respBody))
		fmt.Println("❌", msg)
		http.Error(w, msg, 500)
		return
	}
	fmt.Printf("📤 Anthropic 响应: content=%d items\n", len(anthropicResp.Content))

	openAIResp := convertAnthropicToOpenAI(&anthropicResp)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(openAIResp)
}

// ------------------- 数据结构 -------------------
// OpenAI 的 content 可以是字符串，也可以是数组
type ContentPart struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type OpenAIRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string      `json:"role"`
		Content interface{} `json:"content"` // 字符串或数组
	} `json:"messages"`
	MaxTokens int  `json:"max_tokens"`
	Stream    bool `json:"stream,omitempty"`
}

type AnthropicRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	MaxTokens int  `json:"max_tokens"`
	Stream    bool `json:"stream,omitempty"`
}

type AnthropicResponse struct {
	Content []struct {
		Text string `json:"text"`
	} `json:"content"`
	Usage struct {
		InputTokens  int `json:"input_tokens"`
		OutputTokens int `json:"output_tokens"`
	} `json:"usage"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// ------------------- 转换器 -------------------
// extractContent 从 interface{} 中提取文本内容
func extractContent(content interface{}) string {
	switch v := content.(type) {
	case string:
		return v
	case []interface{}:
		// 数组格式: [{"type": "text", "text": "xxx"}, ...]
		var texts []string
		for _, item := range v {
			if m, ok := item.(map[string]interface{}); ok {
				if m["type"] == "text" {
					if text, ok := m["text"].(string); ok {
						texts = append(texts, text)
					}
				}
			}
		}
		return strings.Join(texts, "\n")
	default:
		return ""
	}
}

func convertOpenAIToAnthropic(req *OpenAIRequest) *AnthropicRequest {
	ar := &AnthropicRequest{
		Model:     req.Model,
		MaxTokens: 4096,
		Stream:    req.Stream,
	}
	for _, msg := range req.Messages {
		ar.Messages = append(ar.Messages, struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			Role:    msg.Role,
			Content: extractContent(msg.Content),
		})
	}
	return ar
}

func convertAnthropicToOpenAI(resp *AnthropicResponse) *OpenAIResponse {
	or := &OpenAIResponse{}
	if len(resp.Content) > 0 {
		or.Choices = append(or.Choices, struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		}{})
		or.Choices[0].Message.Content = resp.Content[0].Text
	}
	or.Usage.PromptTokens = resp.Usage.InputTokens
	or.Usage.CompletionTokens = resp.Usage.OutputTokens
	or.Usage.TotalTokens = resp.Usage.InputTokens + resp.Usage.OutputTokens
	return or
}

// handleStreamResponse 处理流式响应
func handleStreamResponse(w http.ResponseWriter, resp *http.Response) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming not supported", 500)
		return
	}

	scanner := bufio.NewScanner(resp.Body)
	var currentData string
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "data: ") {
			currentData = strings.TrimPrefix(line, "data: ")
		} else if line == "" && currentData != "" {
			// 空行表示一个事件结束
			if currentData == "[DONE]" {
				fmt.Fprintf(w, "data: [DONE]\n\n")
				flusher.Flush()
				break
			}

			var event map[string]interface{}
			if err := json.Unmarshal([]byte(currentData), &event); err != nil {
				currentData = ""
				continue
			}

			openaiChunk := convertAnthropicStreamToOpenAI(event)
			currentData = ""

			if openaiChunk == nil {
				continue
			}

			chunkData, _ := json.Marshal(openaiChunk)
			fmt.Fprintf(w, "data: %s\n\n", chunkData)
			flusher.Flush()
		}
	}
	fmt.Println("📤 流式响应完成")
}

func convertAnthropicStreamToOpenAI(event map[string]interface{}) map[string]interface{} {
	eventType, _ := event["type"].(string)

	switch eventType {
	case "content_block_delta":
		delta, _ := event["delta"].(map[string]interface{})
		// delta 结构: {"type": "text_delta", "text": "xxx"}
		text, _ := delta["text"].(string)
		return map[string]interface{}{
			"id":      "chatcmpl-" + time.Now().Format("20060102150405"),
			"object":  "chat.completion.chunk",
			"created": time.Now().Unix(),
			"model":   "claude",
			"choices": []map[string]interface{}{
				{
					"index": 0,
					"delta": map[string]interface{}{
						"content": text,
					},
				},
			},
		}
	case "message_stop":
		return map[string]interface{}{
			"id":      "chatcmpl-" + time.Now().Format("20060102150405"),
			"object":  "chat.completion.chunk",
			"created": time.Now().Unix(),
			"model":   "claude",
			"choices": []map[string]interface{}{
				{
					"index":          0,
					"finish_reason":  "stop",
					"delta":          map[string]interface{}{},
				},
			},
		}
	}
	return nil
}
