package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func encryptPassword(key []byte, password string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(password))

	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(password))

	return hex.EncodeToString(cipherText), nil
}

func main() {
	var password string
	fmt.Println("Welcome to GoVault")
	fmt.Print("Enter the password: ")
	fmt.Scanln(&password)

	// TODO: Import key from .env
	key := []byte("mysecretkeyashay")

	fmt.Println(encryptPassword(key, password))

}
