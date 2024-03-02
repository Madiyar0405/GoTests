package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFilteredProductsIntegration(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/filtered-products?color=white&sort=name&page=1", nil)
	assert.NoError(t, err, "Error creating request")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Response status should be OK")

}

func setupRouter() http.Handler {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	// Define routes
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/products", func(c *gin.Context) {
		products := getMockProducts()

		c.HTML(http.StatusOK, "products.html", gin.H{
			"Products": products,
		})
	})

	router.GET("/filtered-products", func(c *gin.Context) {

	})

	return router
}
