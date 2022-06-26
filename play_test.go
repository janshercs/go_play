package goplay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringRules(t *testing.T) {
	assert.Equal(t, []string{"a"}, getStringsThatMeetRules([]string{""}, []string{""}))
}
