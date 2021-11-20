package main

import (
	"birthday-telegram-bot/parsing"
	"birthday-telegram-bot/telegram"
	"birthday-telegram-bot/tools"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		msg := `
			First param is location of namedays file in csv format, 
			Second param is location of birthday file in json format
		`
		panic(msg)
	}
	log.Println("first param is", args[0])
	log.Println("second param is", args[1])

	namedays := parsing.ParseNameDays(args[0])
	people := parsing.ParseBirthDays(args[1])

	message := tools.GenerateMessage(namedays, people)

	telegram := telegram.CreateTelegram(os.Getenv("TELEGRAM_API"), os.Getenv("MESSAGE_ID"))
	if len(*message) > 0 {
		telegram.SendMessage(message)
	}
}
