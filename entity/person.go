package entity

import (
	"fmt"
	"strings"
	"time"
)

const (
	timeFormat = "02-01-2006"
)

type Date time.Time

type Person struct {
	FirstName           string `json: "firstname"`
	LastName            string `json: "lastname"`
	NickName            string `json: "nickname"`
	BirthDay            Date   `json: "birthday"`
	NotifyBefore        int    `json: "notifyBefore,omitempty"`
	NotifyBeforeNameDay int    `json: "NotifyBeforeNameDay,omitempty"`
	FictionalName       string
}

// JsonDate deserialization
func (t *Date) UnmarshalJSON(data []byte) (err error) {
	newTime, err := time.ParseInLocation("\""+timeFormat+"\"", string(data), time.Local)
	*t = Date(newTime)
	return
}

// JsonDate serialization
func (t Date) MarshalJSON() ([]byte, error) {
	timeStr := fmt.Sprintf("\"%s\"", time.Time(t).Format(timeFormat))
	return []byte(timeStr), nil
}

// JsonDate serialization
func (t Date) String() string {
	result := fmt.Sprintf("\"%s\"", time.Time(t).Format(timeFormat))
	return strings.Replace(result, "-", "\\-", -1)
}
