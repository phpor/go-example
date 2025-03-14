package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func uploadFile(url string, filePath string, formData map[string]string, username string, password string) error {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// 创建一个缓冲区来存储 multipart/form-data 请求体
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加文件字段到请求体中
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return fmt.Errorf("failed to create form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("failed to copy file content: %v", err)
	}

	// 添加其他表单数据到请求体中
	for key, value := range formData {
		err := writer.WriteField(key, value)
		if err != nil {
			return fmt.Errorf("failed to write form field: %v", err)
		}
	}

	// 关闭 writer，确保所有数据都写入缓冲区
	err = writer.Close()
	if err != nil {
		return fmt.Errorf("failed to close writer: %v", err)
	}

	// 创建 HTTP POST 请求
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// 设置请求头，包括 Content-Type
	req.Header.Set("Content-Type", writer.FormDataContentType())
	// 设置basic认证的header
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
	req.Header.Set("Authorization", basicAuth)

	// 发送 HTTP 请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	// 打印响应状态码和响应体
	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", string(respBody))

	return nil
}

func main() {
	var (
		url      string // 上传文件的 URL
		filePath string
		username string
		password string
	)

	flag.StringVar(&url, "url", os.Getenv("UPLOAD_URL"), "upload url")
	flag.StringVar(&filePath, "file", os.Getenv("FILE_PATH"), "file path")
	flag.StringVar(&username, "username", os.Getenv("USERNAME"), "username")
	flag.StringVar(&password, "password", os.Getenv("PASSWORD"), "password")

	flag.Parse()

	// 其他表单数据
	formData := map[string]string{
		"title": "Sample File",
		"tags":  "document,sample",
	}

	err := uploadFile(url, filePath, formData, username, password)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
