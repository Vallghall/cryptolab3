package main

import (
	"cryptolab3/pkg/des"
	"io/ioutil"
	"log"
)

func main() {
	text, err := ioutil.ReadFile("from.txt")
	if err != nil {
		log.Fatalln(err)
	}
	key, err := ioutil.ReadFile("key.txt")
	if err != nil {
		log.Fatalln(err)
	}
	result, err := des.Encrypt(text, key)
	if err != nil {
		log.Fatalln(err)
	}
	if err = ioutil.WriteFile("to.txt", result, 0755); err != nil {
		log.Fatalln(err)
	}
	decrypted, err := des.Decrypt(result, key)
	if err != nil {
		log.Fatalln(err)
	}
	if err = ioutil.WriteFile("after.txt", decrypted, 0755); err != nil {
		log.Fatalln(err)
	}
}
