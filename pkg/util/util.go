package util

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CreateDataDirIfNotExists() {
	if _, err := os.Stat("data/"); errors.Is(err, os.ErrNotExist) {
		os.Mkdir("data", 0755)
	}
}

// RandomInt returns a random integer between min and max (inclusive)
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString returns a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomEmail returns a random email
func RandomEmail(n int) string {
	return fmt.Sprintf("%s@email.com", RandomString(n))
}
