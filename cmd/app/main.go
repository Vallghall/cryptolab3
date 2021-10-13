package main

import (
	"cryptolab3/pkg/des"
	"log"
	"os"
)

func main() {
	text, err := os.ReadFile("from.txt")
	if err != nil {
		log.Fatalln(err)
	}
	key, err := os.ReadFile("key.txt")
	if err != nil {
		log.Fatalln(err)
	}
	result, err := des.Encrypt(text, key)
	if err != nil {
		log.Fatalln(err)
	}
	if err = os.WriteFile("to.txt", result, 0755); err != nil {
		log.Fatalln(err)
	}
	decrypted, err := des.Decrypt(result, key)
	if err != nil {
		log.Fatalln(err)
	}
	if err = os.WriteFile("after.txt", decrypted, 0755); err != nil {
		log.Fatalln(err)
	}
}
