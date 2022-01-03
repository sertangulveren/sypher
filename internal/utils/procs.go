package utils

import (
	rand2 "math/rand"
)

const KeySize = 32

func GenerateKey() string {
	var collection = []byte("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	buf := make([]byte, KeySize)
	for i := range buf {
		buf[i] = collection[rand2.Intn(len(collection))]
	}
	return string(buf)
}
