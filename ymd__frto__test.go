package cmmn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test__ymd__frto(t *testing.T) {
	fr, _ := Ymd__new("2025-08-16")
	to, _ := Ymd__new("2025-08-31")
	ymd__frto := Ymd__frto__new(*to, *fr)
	assert.Equal(t, "2025-08-16", ymd__frto.Get__fr().Get__yyyy_mm_dd())
	assert.Equal(t, "2025-08-31", ymd__frto.Get__to().Get__yyyy_mm_dd())
	assert.Equal(t, 16, ymd__frto.Get__days())
}
