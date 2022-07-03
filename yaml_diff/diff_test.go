package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	var tests = []struct {
		name     string
		src      interface{}
		dest     interface{}
		expected bool
	}{
		{
			name:     "when source and dest are different kinds, then it should return false",
			src:      []string{},
			dest:     map[string]string{},
			expected: false,
		},
		{
			name:     "when source and dest are same kind but have different values, then it should return false",
			src:      "source",
			dest:     "dest",
			expected: false,
		},
		{
			name:     "when source and dest are same kind and have the same values, then it should return true",
			src:      "source",
			dest:     "source",
			expected: true,
		},
		{
			name:     "when source and dest are maps and have the different number of key values, then it should return false",
			src:      map[string]string{"foo": "bar", "baz": "boi"},
			dest:     map[string]string{"foo": "bar"},
			expected: false,
		},
		{
			name:     "when source and dest are maps and have different keys, then it should return false",
			src:      map[string]string{"baz": "boi"},
			dest:     map[string]string{"foo": "bar"},
			expected: false,
		},
		{
			name:     "when source and dest are maps and have same key values, then it should return true",
			src:      map[string]interface{}{"foo": "bar", "baz": 5, "boz": 2.0},
			dest:     map[string]interface{}{"foo": "bar", "baz": 5, "boz": 2.0},
			expected: true,
		},
		{
			name:     "when source and dest are maps and have same nested key values, then it should return true",
			src:      map[string]interface{}{"foo": "bar", "baz": map[string]interface{}{"bam": 5, "boz": 2.0}},
			dest:     map[string]interface{}{"foo": "bar", "baz": map[string]interface{}{"bam": 5, "boz": 2.0}},
			expected: true,
		},
		{
			name:     "when source and dest are slices and have different values, then it should return false",
			src:      []interface{}{"foo", "bar", "baz", "boi"},
			dest:     []interface{}{"foo", "bar"},
			expected: false,
		},
		{
			name:     "when source and dest are slices and have the same values, then it should return true",
			src:      []interface{}{"foo", 5},
			dest:     []interface{}{"foo", 5},
			expected: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			res, _ := Compare(tt.src, tt.dest)
			assert.Equal(t, tt.expected, res)
		})
	}
}
