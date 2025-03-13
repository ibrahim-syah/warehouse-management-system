package utils

import (
	"math/rand"
	"time"
)

func Addr[T any](t T) *T { return &t }

const DateTimeFormat = "2006-01-02 15:04:05"

func RandomLetterSequence(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}
