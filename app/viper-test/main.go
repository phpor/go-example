package main

import "github.com/spf13/viper"

func main() {
	key := "a"
	viper.Set(key, 1)
	println(viper.GetString(key))
}
