package main

import (
	"fmt"
	"math/rand"
	"time"
)

var alpha = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func divmod(numerator, denominator int64) (quotient, remainder int64) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		var randIdx = rand.Intn(len(alpha))
		b[i] = alpha[randIdx]
	}

	return string(b)
}

func intToString(myInt int64) (result string) {
	var remainder int64
	for myInt > 1 {
		myInt, remainder = divmod(myInt, 62)
		result = string(alpha[remainder]) + result
	}

	return result
}

func epochFromDate() (e int64) {
	t := time.Now().UnixMilli() / 100
	st := time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli() / 100
	return t - st
}

func shrtId(l int) (id string) {
	t := epochFromDate()
	timeId := intToString(t)
	random_len := l - len(timeId)
	random_string := RandStringRunes(random_len)
	return timeId + random_string
}

func main() {
	for i := 1; i < 1000; i++ {
		fmt.Println(shrtId(8))
	}
}
