package cmmn

import (
	"strconv"
	"strings"
	"time"
)

type Ym struct {
	t time.Time
}

func (y *Ym) Add(months int) Ym {
	t := y.t.AddDate(0, months, 0)
	ym := Ym__new__ym(
		t.Year(),
		t.Month(),
	)
	return ym
}

func (y *Ym) Get__month() time.Month {
	return y.t.Month()
}

func (y *Ym) Get__year() int {
	return y.t.Year()
}

func (y *Ym) Get__yyyy_mm() string {
	return y.t.Format("2006-01")
}

func Ym__new(yyyy_mm string) (*Ym, error) {
	words := strings.Split(yyyy_mm, "-")
	if len(words) != 2 {
		return nil, Err__ym__not_yyyy_mm
	}
	year, err := strconv.Atoi(words[0])
	if err != nil {
		return nil, err
	}
	month, err := strconv.Atoi(words[1])
	if err != nil {
		return nil, err
	}
	ym := Ym__new__ym(year, time.Month(month))
	return &ym, nil
}

func Ym__new__ym(year int, month time.Month) Ym {
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return Ym{
		t: t,
	}
}
