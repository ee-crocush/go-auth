//go:build ignore

// Этот модуль предназначен для тестовой генерации секретного ключа.
// Он не предназначен для использования в рабочем коде.
// Скомпилировать и запустить можно с помощью команды:
// make generate-key
package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

// generateKey создает случайный секретный ключ длиной 32 байта
// и кодирует его в формате base64.
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
