package goplay

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

type moo struct {
	dur time.Duration
	zoo int
}

func TestDuration(t *testing.T) {
	m := moo{}
	mType := reflect.TypeOf(m)
	for i := 0; i < mType.NumField(); i++ {
		fieldValue := mType.Field(i)
		fmt.Println(fieldValue.Type.Kind())
	}
}

func TestJSON(t *testing.T) {
	js := `0`
	this := foo{}
	err := json.Unmarshal([]byte(js), &this)
	require.NoError(t, err)
}

func TestScope(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run("", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			res := sleeper(ctx)
			fmt.Printf("%d, %s", i, res)
		})
	}
}

func TestParallelScope(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func TestCallFwF(t *testing.T) {
	CallFwF()
}

func TestCreateCsv(t *testing.T) {
	initCsv()
}

func TestDeadlock(t *testing.T) {
	for i := 0; i < 10; i++ {
		deadlock()
	}
}

// func TestDeadlock2(t *testing.T) {
// 	leaker(1 * time.Second)
// }
