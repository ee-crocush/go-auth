package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func generateKey() string {
	key := make([]byte, 32)

	_, err := rand.Read(key)

	if err != nil {
		log.Fatalf("Ошибка генерации ключа: %v", err)
	}

	return base64.StdEncoding.EncodeToString(key)
}

func main() {
	key := generateKey()

	fmt.Println(key)
}
