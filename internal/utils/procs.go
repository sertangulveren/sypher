package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	rand2 "math/rand"
)

const KeySize = 32

func EncodeBase64(src []byte) []byte {
	data := make([]byte, base64.RawStdEncoding.EncodedLen(len(src)))
	base64.RawStdEncoding.Encode(data, src)
	return data
}

func DecodeBase64(src []byte) []byte {
	data := make([]byte, base64.RawStdEncoding.DecodedLen(len(src)))
	n, err := base64.RawStdEncoding.Decode(data, src)
	PanicWithError(err)

	return data[:n]
}

func Encrypt(key string, src []byte) []byte  {
	cBlock, err := aes.NewCipher([]byte(key))
	PanicWithError(err)

	gcm, err := cipher.NewGCM(cBlock)
	PanicWithError(err)

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	PanicWithError(err)

	return gcm.Seal(nonce, nonce, src, nil)
}

func Decrypt(key string, src []byte) []byte {
	cBlock, err := aes.NewCipher([]byte(key))
	PanicWithError(err)

	gcm, err := cipher.NewGCM(cBlock)
	PanicWithError(err)

	nonceSize := gcm.NonceSize()
	nonce, cipherData := src[:nonceSize], src[nonceSize:]

	decrypted, err := gcm.Open(nil, nonce, cipherData, nil)
	PanicWithError(err)

	return decrypted
}

func GenerateKey() string {
	var collection = []byte("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	buf := make([]byte, KeySize)
	for i := range buf {
		buf[i] = collection[rand2.Intn(len(collection))]
	}
	return string(buf)
}
