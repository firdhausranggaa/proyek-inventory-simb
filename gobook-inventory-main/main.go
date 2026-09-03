package main

import (
	"book-inventory/app"
	"book-inventory/auth"
	"book-inventory/db"
	"book-inventory/middleware"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDB()
	defer database.Close()

	r := gin.Default()

	// 1. Konfigurasi CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	appHandler := app.New(database)
	authHandler := auth.New(database)

	api := r.Group("/api")
	{
		api.POST("/register", authHandler.RegisterHandler)
		api.POST("/login", authHandler.LoginHandler)

		protected := api.Group("/")
		protected.Use(middleware.AuthValid)
		{
			// Akses Member & Admin
			protected.GET("/books", appHandler.GetBooks)
			protected.GET("/books/:id", appHandler.GetBookById)
			protected.POST("/borrow", appHandler.BorrowBook)
			protected.POST("/return/:id", appHandler.ReturnBook)
			
			// Endpoint Baru: Riwayat Peminjaman Pribadi
			protected.GET("/borrowings/me", appHandler.GetMyBorrowings)

			// Akses Khusus Admin (Manipulasi Data)
			adminOnly := protected.Group("/")
			adminOnly.Use(middleware.RequireAdmin)
			{
				adminOnly.POST("/books", appHandler.PostBook)
				adminOnly.PUT("/books/:id", appHandler.PutBook)
				adminOnly.DELETE("/books/:id", appHandler.DeleteBook)
			}
		}
	}

	// 2. Konfigurasi Port Dinamis
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}