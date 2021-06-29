package randdate

import (
	"math/rand"
	"time"
)

func RandDate(minYear, maxYear int) (day, date string) {
	if minYear >= maxYear {
		minYear = 1600
		maxYear = 2200
	}
	min := time.Date(minYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(maxYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	fullDate := time.Unix(sec, 0)
	day = string(fullDate.Format("Monday"))
	date = string(fullDate.Format("02-Jan-2006"))
	return day, date
}
