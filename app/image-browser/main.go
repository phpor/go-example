package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"image-browser/static"
	"net/http"
	"os"
	"regexp"
)

var imgExp = regexp.MustCompile(".*\\.(png|jpg|jpeg)$")

// 写一个webserver
func main() {
	root := "./"
	flag.StringVar(&root, "root", "./", "image root")
	flag.Parse()

	tmpl, err := template.ParseFS(static.HTML, "html/*.html")
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.SetHTMLTemplate(tmpl)
	r.Static("/images/", root)
	//r.Static("/static", "./static")
	//r.StaticFS("/static", gin.Dir("./static", false))
	// 加载f中的image.js文件
	r.StaticFS("/static", http.FS(static.JS))
	//r.GET("/static/", func(context *gin.Context) {
	//	http.FileServer(http.FS(js)).ServeHTTP(context.Writer, context.Request)
	//})

	//r.LoadHTMLFiles("index.html")
	//r.StaticFS("/static", http.Dir("./"))
	r.GET("/", func(c *gin.Context) {
		var dir []string
		var images []string
		path := root
		if p := c.GetString("path"); p != "" {
			path = p
		}
		if path != root {
			dir = append(dir, "../")
		}
		// 获取指定目录下的所有目录
		fileInfo, err := os.ReadDir(path)
		if err != nil {
			fmt.Println(err)
		}

		for _, info := range fileInfo {
			if info.IsDir() {
				dir = append(dir, info.Name())
			} else if imgExp.MatchString(info.Name()) {
				images = append(images, info.Name())
			}
		}

		c.HTML(http.StatusOK, "index.html", gin.H{"dir": dir, "image": images})
	})
	_ = r.Run(":8181")
}
