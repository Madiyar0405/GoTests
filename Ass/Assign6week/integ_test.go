package main

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationProductsHandler(t *testing.T) {
	// Запускаем ваше приложение
	go main()

	// Создаем клиент HTTP
	client := http.Client{}

	// Отправляем GET-запрос на маршрут /products
	resp, err := client.Get("http://localhost:8080/products")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	defer resp.Body.Close()

	// Проверяем код состояния
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Читаем тело ответа
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	// Проверяем, что ответ содержит ожидаемый HTML-контент
	assert.Contains(t, string(body), "Products")
}
