package main

import (
	"net/http"
	"testing"
)

func TestEndToEnd(t *testing.T) {
	// Запускаем ваше приложение
	go main()

	// Ожидаем, пока ваш сервер полностью запустится
	// В данном примере мы просто ждем несколько секунд,
	// вы можете использовать специальный механизм ожидания
	// запуска сервера в реальном приложении.
	// Важно убедиться, что сервер полностью запущен, прежде чем начинать тест.
	// Это можно сделать, например, путем использования WaitGroup.

	// time.Sleep(2 * time.Second)

	// Отправляем GET-запрос на страницу /products
	resp, err := http.Get("http://localhost:8080/products")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем код состояния
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	// Здесь вы можете провести дополнительные проверки,
	// например, проверить, что страница содержит ожидаемый контент.
}
