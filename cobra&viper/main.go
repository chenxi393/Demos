/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func readContent(config *viper.Viper) {
	err := config.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if ok {
			fmt.Println("找不到配置文件")
		} else {
			//可能格式不规范
			fmt.Println("解析配置文件出错：", err)
		}
	}

	user1 := config.GetString("section1.user")
	user2 := config.GetString("section2.user")
	height := config.GetInt32("section.body.height")
	weight := config.GetInt32("section.body.weight")

	fmt.Println(user1, user2, height, weight)

}

func readYaml() {
	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("test")
	config.AddConfigPath("$HOME/All_test/cmd_test")
	readContent(config)
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:")
	})
	viper.WatchConfig()
	fmt.Scan()
}

func main() {
	//cmd.Execute()
	readYaml()
}
