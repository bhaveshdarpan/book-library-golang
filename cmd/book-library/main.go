package main

import (
	"com/github/book-go/internal/controller"
	"com/github/book-go/internal/repository"
	"com/github/book-go/internal/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	_ "github.com/lib/pq"
)

func main() {
	r1 := gin.Default()
	dsn := "postgres://book-bank:bookbank@34.143.145.12:5432/bookbank?sslmode=disable"
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
	schema := `
	    CREATE TABLE IF NOT EXISTS BookBank (
	        id UUID PRIMARY KEY,
	        Book_Title VARCHAR(255),
	        Author VARCHAR(255),
	        Genre VARCHAR(255)
	    );
	`
	db.MustExec(schema)

	// Initialize the repository, service, and controller
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)

	// Define the routes
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

	// r.GET("/", func(c *gin.Context) {
	//     c.JSON(http.StatusOK, gin.H{"message": "pong"})
	// })
	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// gin.SetMode(gin.ReleaseMode)
	// r1.Run(":" + port)
	err = r1.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal().AnErr("Failed to start server", err)
	}

}
