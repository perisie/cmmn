package cmmn

import (
	"time"
)

type Ymd struct {
	t time.Time
}

func (y Ymd) Get__mmm() string {
	mm_airports := map[time.Month]string{
		time.January:   "JAN",
		time.February:  "FEB",
		time.March:     "MAR",
		time.April:     "APR",
		time.May:       "MAY",
		time.June:      "JUN",
		time.July:      "JUL",
		time.August:    "AUG",
		time.September: "SEP",
		time.October:   "OCT",
		time.November:  "NOV",
		time.December:  "DEC",
	}
	return mm_airports[y.t.Month()]
}

func (y Ymd) Get__time() time.Time {
	return y.t
}

func (y Ymd) Get__yyyy_mm_dd() string {
	return y.t.Format(time.DateOnly)
}

func Ymd__new(yyyy_mm_dd string) (*Ymd, error) {
	t, err := time.Parse(time.DateOnly, yyyy_mm_dd)
	if err != nil {
		return nil, err
	}
	return &Ymd{
		t: t,
	}, nil
}

func Ymd__new__unsafe(yyyy_mm_dd string) Ymd {
	ymd, _ := Ymd__new(yyyy_mm_dd)
	return *ymd
}
