package diff

import (
	"io/ioutil"
	"reflect"
)

const (
	Source = "src.yaml"
	Sample = "dest.yaml"
)

func ReadFile(filename string) ([]byte, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func Compare(src, dest interface{}) bool {
	s := reflect.ValueOf(src)
	d := reflect.ValueOf(dest)
	if !isSameKind(s, d) {
		return false
	}

	switch s.Kind() {
	case reflect.Map:
		return compareMaps(s, d)
	case reflect.Slice:
		return compareSlices(s, d)
	default:
		return src == dest // for primitives
	}
	// TODO: Compare structs?
}

func isSameKind(s, d reflect.Value) bool {
	return s.Kind() == d.Kind()
}

func compareMaps(src, dest reflect.Value) bool {
	s := src.MapKeys()
	d := dest.MapKeys()
	if len(s) != len(d) {
		return false
	}

	keys := map[string]bool{}

	for _, key := range s {
		keys[key.String()] = true

	}
	for _, key := range d {
		_, ok := keys[key.String()]
		if !ok {
			return false
		}
		if !Compare(src.MapIndex(key).Interface(), dest.MapIndex(key).Interface()) {
			return false
		}
	}
	return true
}

func compareSlices(src, dest reflect.Value) bool {
	if src.Len() != dest.Len() {
		return false
	}

	for i := 0; i < src.Len(); i++ {
		if !Compare(src.Index(i).Interface(), dest.Index(i).Interface()) {
			return false
		}
	}

	return true
}
