package utils

import (
	"bytes"
	"testing"
)

var (
	contentBase64     = []byte("SGVsbG8gV29ybGQ")
	contentPlain      = []byte("Hello World")
	testData          = []byte("0")
	testEncryptedData = "bac57c92c77132f295e7485ce0b6dc3e04e8e7b7b3a19fb7063ed21a72"
)

const TestAesKey = "12345678901234567890123456789012"

func TestEncodeBase64(t *testing.T) {
	result := EncodeBase64(contentPlain)
	if !bytes.Equal(contentBase64, result) {
		t.Errorf("expected %s, got %s", contentBase64, result)
	}
}

func TestDecodeBase64(t *testing.T) {
	result := DecodeBase64(contentBase64)
	if !bytes.Equal(contentPlain, result) {
		t.Errorf("expected %s, got %s", contentPlain, result)
	}
}

func TestEncrypt(t *testing.T) {
	// no error expected
	_ = Encrypt(TestAesKey, testData)
}

func TestDecrypt(t *testing.T) {
	encrypted := Encrypt(TestAesKey, testData)
	res := Decrypt(TestAesKey, encrypted)

	if !bytes.Equal(testData, res) {
		t.Errorf("expected %x, got %x", testData, res)
	}
}
