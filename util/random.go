package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Run when the first time use the util package
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt func return random number in [min, max]
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString
func RandomString(n int) string {
	var sb strings.Builder
	len := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(len)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner random owner name
func RandomOwner() string {
	return RandomString(10)
}

func RandomMoney() int64 {
	return RandomInt(0, 100)
}

func RandomCurrency() string {
	currencies := [3]string{USD, EUR, CAD}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomOwner())
}
