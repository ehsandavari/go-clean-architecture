package valueObjects

import (
	"fmt"
	"time"
)

type Date struct {
	year  int
	month time.Month
	day   time.Weekday
}

func NewDate(year int, month time.Month, day time.Weekday) Date {
	return Date{
		year:  year,
		month: month,
		day:   day,
	}
}

func (d Date) Year() int {
	return d.year
}

func (d Date) Month() time.Month {
	return d.month
}

func (d Date) Day() time.Weekday {
	return d.day
}

func (d Date) String() string {
	return fmt.Sprintf("%d-%02d-%02d", d.year, d.month, d.day)
}
