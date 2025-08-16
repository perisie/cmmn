package cmmn

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test__ymd(t *testing.T) {
	_, err := Ymd__new("")
	assert.Error(t, err)
	yyyy_mm_dd := "2025-08-16"
	ymd, err := Ymd__new(yyyy_mm_dd)
	assert.NoError(t, err)
	assert.Equal(t, yyyy_mm_dd, ymd.Get__yyyy_mm_dd())
	assert.Equal(t, 2025, ymd.Get__time().Year())
	assert.Equal(t, time.August, ymd.Get__time().Month())
	assert.Equal(t, 16, ymd.Get__time().Day())
	assert.Equal(t, "AUG", ymd.Get__mmm())
}

func Test__ymd__unsafe(t *testing.T) {
	yyyy_mm_dd := "2025-08-16"
	ymd := Ymd__new__unsafe(yyyy_mm_dd)
	assert.Equal(t, yyyy_mm_dd, ymd.Get__yyyy_mm_dd())
}
