package goplay

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomIDString(n int) string {
	b := make([]byte, n)
	rand.New(rand.NewSource(time.Now().UnixNano())).Read(b)
	return fmt.Sprintf("%x", b)[:n]
}

func randomIDStringWithPrefix(prefix ...string) string {
	if len(prefix) == 0 {
		prefix = append(prefix, RandomIDString(8))
	}

	// %.8s means 0 width, up to 8 precision (characters)
	return fmt.Sprintf("%.8s-%.8s", prefix[0], RandomIDString(8))
}
