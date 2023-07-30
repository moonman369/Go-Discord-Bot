package config

import (
	"fmt"
	"os"

	// "github.com/Workiva/go-datastructures/threadsafe/err"
	"github.com/joho/godotenv"
)

var (
	Token     string
	BotPrefix string
	// config    *configStruct
)

// type configStruct struct {
// 	Token     string `json:"Token"`
// 	BotPrefix string `json:"BotPrefix"`
// }

func ReadConfig() error {
	fmt.Println("Reading `config.json` file...")

	// file, err := os.ReadFile("./config.json")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return err
	// }

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// config.Token = os.Getenv("TOKEN")
	// config.BotPrefix = os.Getenv("BOT_PREFIX")

	// fmt.Println(string(file))

	// err = json.Unmarshal(file, &config)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return err
	// }

	// Token = config.Token
	// BotPrefix = config.BotPrefix

	Token = os.Getenv("TOKEN")
	BotPrefix = os.Getenv("BOT_PREFIX")

	return nil
}
