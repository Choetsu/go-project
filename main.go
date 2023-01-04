package main

import (
	"log"
	"os"

	"go-project/broadcast"
	"go-project/handler"
	"go-project/payment"
	"go-project/product"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbUrl := os.Getenv("DB_URL") // Get the DB_URL environment variable
	if dbUrl == "" {             // If the DB_URL environment variable is not set, exit the program
		dbUrl = "test:password@tcp(127.0.0.1:3306)/go-project?charset=utf8&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{}) // Open a connection to the database
	if err != nil {                                         // If there is an error, exit the program
		log.Fatal(err.Error())
	}

	// Migrate the schema for the payment and product models to the database
	db.AutoMigrate(&product.Product{}, &payment.Payment{})
	b := broadcast.NewBroadcaster(20) // Create a new broadcast

	productRepository := product.NewProductRepository(db)          // Create a new product repository
	productService := product.NewProductService(productRepository) // Create a new product service
	productHandler := handler.NewProductHandler(productService)    // Create a new product handler

	paymentRepository := payment.NewPaymentRepository(db)          // Create a new payment repository
	paymentService := payment.NewPaymentService(paymentRepository) // Create a new payment service
	paymentHandler := handler.NewPaymenthandler(paymentService, b) // Create a new payment handler

	r := gin.Default() // Create a new Gin router

	// //Add a route to get the health of the server
	api := r.Group("/api")                            // Create a group of routes for the API
	api.POST("/product", productHandler.Create)       // Add a route to create a new product
	api.GET("/products", productHandler.GetAll)       // Add a route to get all products
	api.GET("/product/:id", productHandler.GetByID)   // Add a route to get a product by its ID
	api.PUT("/product/:id", productHandler.Update)    // Add a route to update a product
	api.DELETE("/product/:id", productHandler.Delete) // Add a route to delete a product

	api.POST("/payment", paymentHandler.Create)       // Add a route to create a new payment
	api.GET("/payments", paymentHandler.GetAll)       // Add a route to get all payments
	api.GET("/payment/:id", paymentHandler.GetByID)   // Add a route to get a payment by its ID
	api.PUT("/payment/:id", paymentHandler.Update)    // Add a route to update a payment
	api.DELETE("/payment/:id", paymentHandler.Delete) // Add a route to delete a payment

	api.GET("/payment/stream", paymentHandler.Stream) // Add a route to delete a payment

	r.Run(":3000") // Start the server on the port specified in the PORT environment variable
}
