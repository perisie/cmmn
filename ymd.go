package cmmn

import (
	"time"
)

type Ymd struct {
	t time.Time
}

func (y Ymd) Get__mmm() string {
	return g__mmm[y.t.Month()]
}

func (y Ymd) Get__time() time.Time {
	return y.t
}

func (y Ymd) Get__yyyy_mm_dd() string {
	return y.t.Format(time.DateOnly)
}

func (y Ymd) Is__between(fr Ymd, to Ymd) bool {
	if fr.Get__time().After(to.Get__time()) {
		tmp := fr
		fr = to
		to = tmp
	}
	t := y.Get__time()
	if t.Year() < fr.Get__time().Year() {
		return false
	}
	if t.Month() < fr.Get__time().Month() {
		return false
	}
	if t.Day() < fr.Get__time().Day() {
		return false
	}
	if t.Year() > to.Get__time().Year() {
		return false
	}
	if t.Month() > to.Get__time().Month() {
		return false
	}
	if t.Day() > to.Get__time().Day() {
		return false
	}
	return true
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
