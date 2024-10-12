package crud

import (
	"time"

	"golang.org/x/exp/rand"
)

// const charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+,.?/:;{}[]`~"
const charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GeneratePassword(length int) string {
	r := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[r.Intn(len(charset))]
	}

	return string(password)
}
