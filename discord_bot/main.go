package main

import (
	"fmt"
	"go_learn/discord_bot/bot"
	"go_learn/discord_bot/config"
)

func main(){

	err:=config.ReadConfig()

	if err != nil {

		fmt.Println(err.Error())
	}

	bot.Start()

	<-make(chan struct{})
	return

}