/*
TODO: change to https://github.com/go-faker/faker
*/

package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	// in fact return a number in[min, min+(max-min)]
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	// Create random string
	for i := 0; i < n; i++ {
		char := alphabet[rand.Intn(k)]
		sb.WriteByte(char)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomAmount() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
