package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ProtocHandlerFunc(w http.ResponseWriter, req *http.Request) {
	workDir := "/tmp/proto"
	//tmpDir := filepath.Join(workDir, "proto-1")
	tmpDir, err := os.MkdirTemp(workDir, "proto-*")
	if err != nil {
		log.Println("Create temp dir:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()
	// curl -v -F "file=@test.txt" http://localhost:8181/upload
	file, header, err := req.FormFile("file")

	if file == nil {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    0,
			"message": "no upload file",
		})
		return
	}

	if err != nil {
		log.Println("Parse form file:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		_ = file.Close()
		_ = req.MultipartForm.RemoveAll() // Seen from go source code, req.MultipartForm not nil after call FormFile(..)
	}()

	filename := req.FormValue("filename")
	log.Println("filename FromValue: " + filename)
	if filename == "" {
		filename = header.Filename // 只能取到文件名，无法取到文件路径，即使客户端传了文件路径，golang也会给去掉
		log.Println("filename FromHeader: " + filename)
	}
	dirName := filepath.Base(filepath.Dir(filename))
	filename = filepath.Base(filename)

	if err := os.MkdirAll(filepath.Join(tmpDir, dirName), 0755); err != nil {
		log.Println("Create dir fail:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dstPath := filepath.Join(tmpDir, dirName, filename)
	log.Println("dstPath: ", dstPath)

	// Large file (>32MB) will store in tmp directory
	// The quickest operation is call os.Move instead of os.Copy
	// Note: it seems not working well, os.Rename might be failed

	var copyErr error
	// if osFile, ok := file.(*os.File); ok && fileExists(osFile.Name()) {
	// 	tmpUploadPath := osFile.Name()
	// 	osFile.Close() // Windows can not rename opened file
	// 	log.Printf("Move %s -> %s", tmpUploadPath, dstPath)
	// 	copyErr = os.Rename(tmpUploadPath, dstPath)
	// } else {
	dst, err := os.Create(dstPath)
	if err != nil {
		log.Println("Create file:", err)
		http.Error(w, "File create "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Note: very large size file might cause poor performance
	// _, copyErr = io.Copy(dst, file)
	buf := make([]byte, 8192)
	_, copyErr = io.CopyBuffer(dst, file, buf)
	_ = dst.Close()
	if copyErr != nil {
		log.Println("Handle upload file:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 执行一条命令  docker run --rm
	args := []string{"run", "--rm", "-v", tmpDir + ":" + tmpDir, "-i", "-w", tmpDir, "registry.api.weibo.com/passport/protobuf:v2.0.2", "bash", "-x"}
	cmd := exec.Command("docker", args...)
	// 给cmd写入输入流

	// 读取并打印cmd的标准输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	relativePath := filepath.Join(dirName, filename)
	cmd.Stdin = strings.NewReader(protoc(relativePath))
	log.Println(cmd.String())
	err = cmd.Run()
	if err != nil {
		log.Println("exec command:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	CompressToZip(w, filepath.Dir(dstPath))

	//w.Header().Set("Content-Type", "application/json;charset=utf-8")
	//
	//json.NewEncoder(w).Encode(map[string]interface{}{
	//	"code":    0,
	//	"message": "upload success",
	//})
}

func protoc(filename string) string {
	return `protoc  -I. -I/root/go/pkg/mod/github.com/googleapis/googleapis@v0.0.0-20240531082521-2244fe420817/\
	--plugin=protoc-gen-grpc-gateway:. \
	--plugin=protoc-gen-go:. \
	--plugin=protoc-gen-go-grpc:. \
	--plugin=protoc-gen-openapiv2:. \
	--grpc-gateway_out . \
	--go-grpc_out . \
	--go_out . \
	--go_opt=paths=source_relative \
	--grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	--openapiv2_out=. \
	--openapiv2_opt=logtostderr=true ` + filename
}
