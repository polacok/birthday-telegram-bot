package parsing

import (
	"birthday-telegram-bot/entity"
	"encoding/json"
	"io/ioutil"
	"log"
)

type BirthDays struct {
	People []entity.Person `json: "people"`
}

func ParseBirthDays(fileLocalion string) []*entity.Person {
	data, err := ioutil.ReadFile(fileLocalion)
	if err != nil {
		log.Println("Invalid birthday file")
		panic(err)
	}
	var birthdays BirthDays
	err = json.Unmarshal(data, &birthdays)
	if err != nil {
		log.Println("Invalid structure of birthday file")
		panic(err)
	}
	result := make([]*entity.Person, len(birthdays.People))
	for index, _ := range birthdays.People {
		// musime ist cez indexy a nie cez iterator
		result[index] = &birthdays.People[index]
	}
	return result
}
