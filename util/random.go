package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomNumber(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		sb.WriteByte(alphabet[rand.Intn(k)])
	}
	return sb.String()
}

func RandomOwner() string {
	return randomString(6)
}

func RandomMoney() int64 {
	return randomNumber(1, 1000)
}
func RandomCurrency() string {
	currencies := []string{"EUR", "IDR", "USD", "CAD"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}
