package main

import (
	"fmt"
	"log"
	"net/http"

	// "log"
	// "net/http"

	"github.com/moonman369/Go-Discord-Bot/bot"
	"github.com/moonman369/Go-Discord-Bot/config"

	// "github.com/moonman369/Go-Discord-Bot/gpt"
	"github.com/gorilla/mux"
)

func main() {
	// gpt.SendPrompt("Hello")
	r := mux.NewRouter()

	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})

	// <-make(chan struct{})

	return
}
