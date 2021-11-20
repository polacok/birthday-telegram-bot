package tools

import (
	"birthday-telegram-bot/entity"
	"bytes"
	"fmt"
	"log"
	"time"
)

func DateToIntMapping(time *time.Time) int {
	return int(time.Month())<<9 + time.Day()
}

func GenerateMessage(nameDays map[time.Time][]string, people []*entity.Person) *string {
	nowTime := time.Now()
	indexInMaps := DateToIntMapping(&nowTime)

	nameToTimeMap := nameToTimeMap(nameDays)
	timeDayMonthToNamesMap := timeDayMonthToNamesMap(nameDays)
	fillFictionalName(people, nameToTimeMap, timeDayMonthToNamesMap)

	names := groupByFictionalName(people)
	var resultBuffer bytes.Buffer

	var nameDayBuffer bytes.Buffer
	if nameday, ok := timeDayMonthToNamesMap[indexInMaps]; ok {
		for _, firstName := range nameday {
			if people, ok := names[firstName]; ok {
				for _, person := range people {
					printPersonToBuffer(&nameDayBuffer, person, person.NotifyBeforeNameDay)
				}
			}
		}
	} else {
		log.Fatalf("Someone always has namedays")
	}

	if nameDayBuffer.Len() != 0 {
		resultBuffer.WriteString("*Meniny*\n")
		resultBuffer.Write(nameDayBuffer.Bytes())
	}

	birthDays := convertPeopleToBirthDayMapWithNegativeOffset(people)
	var birthDayBuffer bytes.Buffer
	if people, ok := birthDays[indexInMaps]; ok {
		for _, person := range people {
			printPersonToBuffer(&birthDayBuffer, person, person.NotifyBefore)
		}
	}

	if birthDayBuffer.Len() != 0 {
		if resultBuffer.Len() != 0 {
			resultBuffer.WriteString("\n")
		}
		resultBuffer.WriteString("*Narodeniny*\n")
		resultBuffer.Write(birthDayBuffer.Bytes())
	}
	result := resultBuffer.String()
	return &result
}

func nameToTimeMap(nameDays map[time.Time][]string) map[string]time.Time {
	result := make(map[string]time.Time)
	for time, firstNames := range nameDays {
		for _, firstName := range firstNames {
			result[firstName] = time
		}
	}
	return result
}

func timeDayMonthToNamesMap(nameDays map[time.Time][]string) map[int][]string {
	result := make(map[int][]string)
	for time, firstNames := range nameDays {
		result[DateToIntMapping(&time)] = firstNames
	}
	return result
}

func printPersonToBuffer(buffer *bytes.Buffer, person *entity.Person, dayOffset int) {
	buffer.WriteString(fmt.Sprintf("%s \\(%s %s\\) \\- %s in %d days\n",
		person.NickName, person.FirstName, person.LastName, person.BirthDay.String(), dayOffset))
}

func fillFictionalName(people []*entity.Person, nameToNamedayMap map[string]time.Time, nameDaysToNamesList map[int][]string) {
	for _, person := range people {
		originalNameday := nameToNamedayMap[person.FirstName]
		fictionalNameDay := originalNameday.AddDate(0, 0, -1*person.NotifyBeforeNameDay)
		if fictionalNames, ok := nameDaysToNamesList[DateToIntMapping(&fictionalNameDay)]; ok {
			person.FictionalName = fictionalNames[0]
		} else {
			person.FictionalName = person.FirstName
		}
	}
}

func convertPeopleToBirthDayMapWithNegativeOffset(people []*entity.Person) map[int][]*entity.Person {
	result := make(map[int][]*entity.Person)
	for _, person := range people {
		notifyTime := time.Time(person.BirthDay).AddDate(0, 0, -1*person.NotifyBefore)
		index := DateToIntMapping(&notifyTime)
		if personWithSameDay, ok := result[index]; ok {
			personWithSameDay = append(personWithSameDay, person)
		} else {
			tmp := make([]*entity.Person, 1)
			tmp[0] = person
			result[index] = tmp
		}
	}
	return result
}

func groupByFictionalName(people []*entity.Person) map[string][]*entity.Person {
	result := make(map[string][]*entity.Person)
	for _, person := range people {
		personWithSameDay, ok := result[person.FictionalName]
		if ok {
			personWithSameDay = append(personWithSameDay, person)
		} else {
			tmp := make([]*entity.Person, 1)
			tmp[0] = person
			result[person.FictionalName] = tmp
		}
	}
	return result
}
