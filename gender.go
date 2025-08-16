package cmmn

import "strings"

type Gender string

const (
	Gender__female Gender = "female"
	Gender__male   Gender = "male"
)

func Gender__new(s string) Gender {
	gender := Gender__female
	if strings.ToLower(s) == string(Gender__male) {
		gender = Gender__male
	}
	return gender
}
