package main

import (
	"com/github/book-go/internal/controller"
	"com/github/book-go/internal/repository"
	"com/github/book-go/internal/service"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	_ "github.com/lib/pq"
)

func main() {
	// Create a new Gin router
	r1 := gin.Default()

	// Set up CORS with default settings
	r1.Use(cors.Default())

	// Database connection
	dsn := "postgres://book-bank:bookbank@10.31.0.6:5432/bookbank?sslmode=disable"
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal().AnErr("DB not started properly", err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatal().AnErr("Cannot connect to the database: %v", err)
	}
	log.Print("Connected to the database")

	// Database schema
	schema := `
	    CREATE TABLE IF NOT EXISTS BookBank (
	        id UUID PRIMARY KEY,
	        Book_Title VARCHAR(255),
	        Author VARCHAR(255),
	        Genre VARCHAR(255)
	    );
	`
	db.MustExec(schema)

	// Initialize repository, service, and controller
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)

	// Define routes
	r1.GET("/healthz", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"status": "Healthy",
		})
	})
	
	r1.GET("/api/books", func(c *gin.Context) {
		bookController.GetAllBooks(c.Writer, c.Request)
	})
	r1.GET("/api/books/:id", func(c *gin.Context) {
		bookController.GetBook(c.Writer, c.Request)
	})
	r1.POST("/api/books", func(c *gin.Context) {
		bookController.AddBook(c.Writer, c.Request)
	})
	r1.PUT("/api/books/:id", func(c *gin.Context) {
		bookController.UpdateBook(c.Writer, c.Request)
	})
	r1.DELETE("/api/books/:id", func(c *gin.Context) {
		bookController.DeleteBook(c.Writer, c.Request)
	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = r1.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal().AnErr("Failed to start server", err)
	}
}








// package main

// import (
// 	"com/github/book-go/internal/controller"
// 	"com/github/book-go/internal/repository"
// 	"com/github/book-go/internal/service"
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// 	"github.com/jmoiron/sqlx"
// 	"github.com/rs/zerolog/log"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	// Create a new Gin router
// 	r1 := gin.Default()

// 	// Set up CORS
// 	corsConfig := cors.Config{
// 		AllowOrigins: []string{"http://10.148.0.9:3000"}, // Allow your React app
// 		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
// 		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
// 	}
// 	r1.Use(cors.New(corsConfig))

// 	// Database connection
// 	dsn := "postgres://book-bank:bookbank@10.31.0.6:5432/bookbank?sslmode=disable"
// 	db, err := sqlx.Open("postgres", dsn)
// 	if err != nil {
// 		log.Fatal().AnErr("DB not started properly", err)
// 		return
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal().AnErr("Cannot connect to the database: %v", err)
// 	}
// 	log.Print("Connected to the database")

// 	// Database schema
// 	schema := `
// 	    CREATE TABLE IF NOT EXISTS BookBank (
// 	        id UUID PRIMARY KEY,
// 	        Book_Title VARCHAR(255),
// 	        Author VARCHAR(255),
// 	        Genre VARCHAR(255)
// 	    );
// 	`
// 	db.MustExec(schema)

// 	// Initialize repository, service, and controller
// 	bookRepository := repository.NewBookRepository(db)
// 	bookService := service.NewBookService(bookRepository)
// 	bookController := controller.NewBookController(bookService)

// 	// Define routes
// 	r1.GET("/healthz", func(c *gin.Context){
// 		c.JSON(http.StatusOK, gin.H{
// 			"status": "Healthy",
// 		})
// 	})
	
// 	r1.GET("/api/books", func(c *gin.Context) {
// 		bookController.GetAllBooks(c.Writer, c.Request)
// 	})
// 	r1.GET("/api/books/:id", func(c *gin.Context) {
// 		bookController.GetBook(c.Writer, c.Request)
// 	})
// 	r1.POST("/api/books", func(c *gin.Context) {
// 		bookController.AddBook(c.Writer, c.Request)
// 	})
// 	r1.PUT("/api/books/:id", func(c *gin.Context) {
// 		bookController.UpdateBook(c.Writer, c.Request)
// 	})
// 	r1.DELETE("/api/books/:id", func(c *gin.Context) {
// 		bookController.DeleteBook(c.Writer, c.Request)
// 	})

// 	// Start the server
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 	}
// 	err = r1.Run(fmt.Sprintf(":%s", port))
// 	if err != nil {
// 		log.Fatal().AnErr("Failed to start server", err)
// 	}
// }
