package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	msg := "This is some text to be sent over an http connection!!"

	encoded := encode(msg)
	fmt.Println("base64 encoded msg: ", encoded)

	s, err := decode(encoded)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("decoded base64: ", s)
}

func encode(msg string) string {
	return base64.URLEncoding.EncodeToString([]byte(msg))
}

func decode(encoded string) (string, error) {
	s, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return "", fmt.Errorf("couldn't decode string %w", err)
	}
	return string(s), nil
}
