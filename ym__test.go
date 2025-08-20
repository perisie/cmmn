package cmmn

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test__ym(t *testing.T) {
	yyyy_mm := "2025-07"
	ym, err := Ym__new(yyyy_mm)
	require.NoError(t, err)
	require.Equal(t, yyyy_mm, ym.Get__yyyy_mm())
	require.Equal(t, 2025, ym.Get__year())
	require.Equal(t, time.July, ym.Get__month())
	ym__prev := ym.Add(-7)
	require.Equal(t, "2024-12", ym__prev.Get__yyyy_mm())
	ym__next := ym.Add(6)
	require.Equal(t, "2026-01", ym__next.Get__yyyy_mm())
}

func Test__ym__bad(t *testing.T) {
	_, err := Ym__new("")
	require.Error(t, err)
	_, err = Ym__new("-")
	require.Error(t, err)
	_, err = Ym__new("2025-")
	require.Error(t, err)
}
