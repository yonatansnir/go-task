package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomID() string {
	seconds := time.Now().UnixMilli() / 1000
	hex := strconv.FormatInt(seconds, 16)

	char := "abcdef0123456789"
	str := ""
	for i := 0; i < 16; i++ {
		index := rand.Intn(len([]rune(char)))
		str += string(char[index])
	}

	return hex + str
}