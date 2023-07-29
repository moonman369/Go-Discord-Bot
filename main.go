package main

import (
	"fmt"

	"github.com/moonman369/Go-Discord-Bot/bot"
	"github.com/moonman369/Go-Discord-Bot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()
	<-make(chan struct{})
	return
}
