package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFilterProductsByColor(t *testing.T) {
	products := getMockProducts()

	filteredProducts := filterProductsByColor(products, "white")
	assert.Equal(t, 3, len(filteredProducts), "Filtered products count should be 3 for color white")

	filteredProducts = filterProductsByColor(products, "black")
	assert.Equal(t, 3, len(filteredProducts), "Filtered products count should be 3 for color black")

	filteredProducts = filterProductsByColor(products, "red")
	assert.Equal(t, 0, len(filteredProducts), "Filtered products count should be 0 for color red")
}

func TestSortProducts(t *testing.T) {
	products := getMockProducts()

	sortedProducts := sortProducts(products, "name")
	assert.Equal(t, "Jacket", sortedProducts[0].Name, "First product should be Jacket when sorted by name")

	sortedProducts = sortProducts(products, "price")
	assert.Equal(t, int16(5000), sortedProducts[0].Price, "First product price should be 5000 when sorted by price")
}

func TestPaginateProducts(t *testing.T) {
	products := getMockProducts()

	paginatedProducts := paginateProducts(products, "1", 2)
	assert.Equal(t, 2, len(paginatedProducts), "Paginated products count should be 2 for page 1")

	paginatedProducts = paginateProducts(products, "2", 2)
	assert.Equal(t, 2, len(paginatedProducts), "Paginated products count should be 2 for page 2")

	paginatedProducts = paginateProducts(products, "3", 2)
	assert.Equal(t, 2, len(paginatedProducts), "Paginated products count should be 2 for page 3")

	paginatedProducts = paginateProducts(products, "4", 2)
	assert.Equal(t, 0, len(paginatedProducts), "Paginated products count should be 0 for page 4")
}

func TestCheckError(t *testing.T) {
	assert.NoError(t, checkError(nil, []Product{{Name: "Test"}}), "No error should be returned for non-empty slice")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	err := checkError(ctx, []Product{})
	assert.Error(t, err, "Error should be returned for empty slice")
	assert.Equal(t, 404, w.Result().StatusCode, "Status code should be 404 for empty slice")
}
