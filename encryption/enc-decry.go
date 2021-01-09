package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// also check 0.19.0

func main() {
	msg := "This is a super secret message which i'll send over http, hopefully!"

	password := "secretpassword"
	// first 16 bytes of bcrypt'd password as the key for AES-128
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatalln("couldn't bcrypt password", err)
	}

	bs = bs[:16]

	result, err := enDecode(bs, msg)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("raw bytes before base64: ", string(result))

	result2, err := enDecode(bs, string(result))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(result2))
}

func enDecode(key []byte, input string) ([]byte, error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("couldn't newCipher %w", err)
	}

	iv := make([]byte, aes.BlockSize)

	s := cipher.NewCTR(b, iv)

	buff := &bytes.Buffer{}
	sw := cipher.StreamWriter{
		S: s,
		W: buff,
	}
	_, err = sw.Write([]byte(input))
	if err != nil {
		return nil, fmt.Errorf("couldn't sw.Write to streamwriter %w", err)
	}

	return buff.Bytes(), nil
}
