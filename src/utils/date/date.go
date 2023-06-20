package date

import (
	"time"
)

const (
	dateFormat = "2006-01-02"
)

type Date struct {
	time time.Time
}

// Now return Date object with current datetime
func Now() *Date {
	return &Date{time: time.Now()}
}

func (d *Date) FormatDate() string {
	return d.time.Format(dateFormat)
}
