package main

import (
	"cryptolab3/pkg/des"
	"fmt"
	"log"
)

func main() {
	text := "some_exampletext"
	key := "abcdabcd"
	result, err := des.Encrypt([]byte(text), []byte(key))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
	decrypted, _ := des.Decrypt(result, []byte(key))
	fmt.Println(string(decrypted))
}
