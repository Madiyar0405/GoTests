package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 3) // Rate limit of 1 request

type Product struct {
	Name        string
	Price       int16
	Color       string
	Description string
	ImagePath   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(gin.DefaultWriter)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

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
		colorFilter := c.Query("color")
		sortBy := c.Query("sort")
		pageStr := c.Query("page")
		itemsPerPage := 3

		products := getMockProducts()

		filteredProducts := filterProductsByColor(products, colorFilter)
		if err := checkError(c, filteredProducts); err != nil {
			return
		}

		sortedProducts := sortProducts(filteredProducts, sortBy)
		if err := checkError(c, sortedProducts); err != nil {
			return
		}

		paginatedProducts := paginateProducts(sortedProducts, pageStr, itemsPerPage)
		if err := checkError(c, paginatedProducts); err != nil {
			return
		}

		// Logging filtered and paginated products
		log.WithFields(logrus.Fields{
			"ColorFilter":       colorFilter,
			"SortBy":            sortBy,
			"Page":              pageStr,
			"ItemsPerPage":      itemsPerPage,
			"FilteredProducts":  len(filteredProducts),
			"PaginatedProducts": len(paginatedProducts),
		}).Info("Filtered and Paginated Products")

		// Sending HTML response to the client
		c.HTML(http.StatusOK, "products.html", gin.H{
			"Products": paginatedProducts,
		})
	})

	//Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	srv := &http.Server{Addr: ":8080", Handler: router}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	log.Println("Server exiting")

	//Shutdown
	<-quit
	log.Println("Shutdown signal received, shutting down gracefully...")
	if err := srv.Shutdown(nil); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}
	log.Println("Server gracefully stopped")
}

func getMockProducts() []Product {
	// Mock data, no error to handle here
	return []Product{
		{"T-shirt", 5000, "white", "Comfortable everyday white T-shirt", "https://images.satu.kz/72306194_futbolka-sols-imperial.jpg"},
		{"T-shirt", 5500, "black", "Comfortable everyday black T-shirt", "https://images.satu.kz/72305201_w600_h600_72305201.jpg"},
		{"Jeans", 10000, "white", "Stylish white jeans for your look", "https://imgcdn.loverepublic.ru/upload/images/22554/2255436764_1_5.jpg"},
		{"Jeans", 6000, "black", "Stylish black jeans for your look", "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ-fRBJWhu5Q5GTJBq6Ki7RvR5Hzo6vDrC16w&usqp=CAU"},
		{"Jacket", 12000, "white", "Warm jacket for cold days", "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSiab0q1ZL5fW_6XKIEvNJ-fAcoyufnBmG7Hg&usqp=CAU"},
		{"Jacket", 9000, "black", "Warm jacket for cold days", "https://momsbox.kz/upload/iblock/6b7/jy8tibyta1vv7fpwr9f7s7qghaxhzygg/%D0%91%D0%B5%D0%B7-%D0%B8%D0%BC%D0%B5%D0%BD%D0%B8-39.jpg"},
	}
}

func filterProductsByColor(products []Product, color string) []Product {
	if color == "" {
		return products
	}

	var filteredProducts []Product
	for _, p := range products {
		if p.Color == color {
			filteredProducts = append(filteredProducts, p)
		}
	}

	return filteredProducts
}

func sortProducts(products []Product, sortBy string) []Product {
	switch sortBy {
	case "name":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Name < products[j].Name
		})
	case "price":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price < products[j].Price
		})
	}
	return products
}

func paginateProducts(products []Product, pageStr string, itemsPerPage int) []Product {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	startIdx := (page - 1) * itemsPerPage
	endIdx := startIdx + itemsPerPage

	if startIdx >= len(products) {
		return []Product{}
	}

	if endIdx > len(products) {
		endIdx = len(products)
	}

	return products[startIdx:endIdx]
}

// checkError is a utility function to check if a slice is empty and send an error response
func checkError(c *gin.Context, data interface{}) error {
	if len(data.([]Product)) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No data available"})
		return fmt.Errorf("no data available")
	}
	return nil
}
