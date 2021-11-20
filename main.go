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
	for _, value := range os.Environ() {
		log.Println(value)
	}
	log.Println("who to trust:", os.Getenv("WHO_TO_TRUST"))
	if len(args) == 1 {
		log.Println(args[0])
	}

	if len(args) != 4 {
		msg := `
			Please provide four params.
			First param is location of namedays file in csv format, 
			Second param is location of birthday file in json format
			Third param is secret telegram token
			Forth param is communicator id to whom send message
		`
		panic(msg)
	}
	log.Println("first param is", args[0])
	log.Println("second param is", args[1])

	namedays := parsing.ParseNameDays(args[0])
	people := parsing.ParseBirthDays(args[1])

	message := tools.GenerateMessage(namedays, people)

	telegram := telegram.CreateTelegram(args[2], args[3])
	if len(*message) > 0 {
		telegram.SendMessage(message)
	}
}
