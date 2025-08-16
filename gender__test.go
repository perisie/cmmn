package cmmn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test__gender(t *testing.T) {
	gender__m := Gender__new("male")
	assert.Equal(t, Gender__male, gender__m)
	gender__f := Gender__new("female")
	assert.Equal(t, Gender__female, gender__f)
}
