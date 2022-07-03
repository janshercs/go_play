package diff

import (
	"fmt"
	"io/ioutil"
	"reflect"
)

type changeType int64

const (
	Source = "src.yaml"
	Sample = "dest.yaml"

	srcOnlyKey changeType = iota
	destOnlyKey
	valueChange
)

// missing/extra key, different value
type Difference struct {
	changeType changeType
	src        interface{}
	dest       interface{}
	key        string
}

func (d Difference) String() string {
	switch d.changeType {
	case valueChange:
		return fmt.Sprintf("values are different: src: %s, dest: %s", d.src, d.dest)
	case srcOnlyKey:
		return fmt.Sprintf("key %s in src, not in dest", d.key)
	default:
		return fmt.Sprintf("key %s in dest, not in src", d.key)
	}
}

func ReadFile(filename string) ([]byte, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func Compare(src, dest interface{}) (bool, []Difference) {
	s := reflect.ValueOf(src)
	d := reflect.ValueOf(dest)
	if !isSameKind(s, d) {
		return false, []Difference{{
			changeType: valueChange,
			src:        src,
			dest:       dest,
		}}
	}

	switch s.Kind() {
	case reflect.Map:
		return compareMaps(s, d)
	case reflect.Slice:
		return compareSlices(s, d)
	default: // Compare primitives
		if src == dest {
			return true, []Difference{}
		}
		return false, []Difference{{
			changeType: valueChange,
			src:        src,
			dest:       dest,
		}}
	}
	// TODO: Compare structs?
}

func isSameKind(s, d reflect.Value) bool {
	return s.Kind() == d.Kind()
}

func compareMaps(src, dest reflect.Value) (bool, []Difference) {
	result := true
	differences := []Difference{}
	s := src.MapKeys()
	d := dest.MapKeys()

	if len(s) != len(d) {
		result = false
	}

	srcKeys := map[string]bool{}

	for _, key := range s {
		srcKeys[key.String()] = true

	}

	for _, key := range d {
		_, ok := srcKeys[key.String()]
		if !ok {
			result = false
			differences = append(differences, Difference{
				changeType: destOnlyKey,
				key:        key.String(),
			})
			continue
		}

		if sameValue, diffs := Compare(src.MapIndex(key).Interface(), dest.MapIndex(key).Interface()); !sameValue {
			result = false
			differences = append(differences, diffs...)
		}
	}

	destKeys := map[string]bool{}

	for _, key := range d {
		destKeys[key.String()] = true

	}

	for _, key := range s {
		_, ok := destKeys[key.String()]
		if !ok {
			result = false
			differences = append(differences, Difference{
				changeType: srcOnlyKey,
				key:        key.String(),
			})
		}
	}

	return result, differences
}

func compareSlices(src, dest reflect.Value) (bool, []Difference) {
	result := true
	differences := []Difference{}
	if src.Len() != dest.Len() {
		result = false
	}

	shorterSliceLength := minInt(src.Len(), dest.Len())
	for i := 0; i < shorterSliceLength; i++ {
		if res, diffs := Compare(src.Index(i).Interface(), dest.Index(i).Interface()); res != true {
			result = false
			differences = append(differences, diffs...)
		}
	}

	return result, differences
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
