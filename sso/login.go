package sso

import (
	"net/http"
	"io/ioutil"
	"net/url"
)

func Login(username string, password string) (uid string, err error) {
	form := make(url.Values)
	form.Set("username", username)
	form.Set("password", password)
	form.Set("entry", "sso")
	form.Set("returntype", "TEXT2")



	response, err := http.PostForm("https://login.sina.com.cn/sso/login.php", form)
	if err != nil {
		return
	}
	defer response.Body.Close()
	content,err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	return string(content),err
}
