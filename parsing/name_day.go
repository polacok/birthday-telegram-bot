package parsing

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

const (
	layout = "01-02"
)

type NameDays struct {
	day map[time.Time][]string
}

func ParseNameDays(fileLocalion string) map[time.Time][]string {
	data, err := ioutil.ReadFile(fileLocalion)
	if err != nil {
		log.Println("Invalid nameday file")
		panic(err)
	}
	result := make(map[time.Time][]string)
	for _, item := range strings.Split(string(data), "\n") {
		tmp := strings.Split(item, ";")
		if len(tmp) != 2 {
			panic(errors.New("Invalid structure of nameday file"))
		}
		date, err := parseTime(tmp[0])
		if err != nil {
			panic(errors.New("Invalid structure of nameday file"))
		}
		result[date] = strings.Split(tmp[1], ",")
	}
	return result
}

func parseTime(value string) (time.Time, error) {
	return time.Parse(layout, value)
}
