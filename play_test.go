package goplay

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	answers = []string{
		"452626", "159452", "52450", "452778", "556208",
	}
)

func TestBruteforce(t *testing.T) {
	assert.Equal(t, answers, bruteForceBinsRules(rules, testcase))
}

func TestMapAlgo(t *testing.T) {
	assert.Equal(t, answers, mapBinsRules(rules, testcase))
}

// cpu: Intel(R) Core(TM) i5-10500 CPU @ 3.10GHz
// BenchmarkBruteforce-12           5484798               213.9 ns/op
// PASS
// ok      go-play 1.857s

func BenchmarkBruteforce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bruteForceBinsRules(rules, testcase)
	}
}

func BenchmarkMapAlgo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapBinsRules(rules, testcase)
	}
}

func TestFoo(t *testing.T) {
	this := foo{
		Bar: "",
		Baz: 30,
	}
	js := `{"bar": "james"}`
	json.Unmarshal([]byte(js), &this) // Unmarshal only populates the fields that are present in bytes
	fmt.Println(this)

	nextjs := `{"baz": 39}`
	json.Unmarshal([]byte(nextjs), &this)
	fmt.Println(this)
}
