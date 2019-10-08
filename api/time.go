package api

import (
	"errors"
	"log"
	"time"
)

// StartDate create a time to start searching in UTC +07.
func StartDate(s string) time.Time {
	const shortForm = "2006-01-02"
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Println("cannot load location Asia/Bangkok.")
	}

	t, err := time.Parse(shortForm, s)
	if err != nil {
		log.Println("cannot convert the date")
	}

	t = t.In(loc)
	t = t.Add(-7 * time.Hour)

	return t
}

// EndDate create a time to end the search.
func EndDate(s string) time.Time {
	t := StartDate(s)
	t = t.AddDate(0, 0, 1)

	return t
}

// ParseDate convert value from characters to date ISO.
func ParseDate(s string) (time.Time, error) {
	const shortForm = "2006-01-02"
	bd, err := time.Parse(shortForm, s)
	if err != nil {
		return time.Time{}, errors.New("invalid date format.")
	}

	return bd, nil
}
