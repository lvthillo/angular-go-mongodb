package main

import (
	"context"
	"movie-app/db"
	"movie-app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	// Create MongoDB Client
	db.Connect()
	// Disconnect client when main function stops running
	defer db.Client.Disconnect(context.Background())

	r := gin.Default()
	r.Use(CORSMiddleware())

	// Routes
	r.GET("/movie", handlers.GetAllMoviesHandler)
	//r.GET("/movie/:id", handlers.GetMovieByIDHandler)
	r.POST("/movie", handlers.AddMovieHandler)
	r.DELETE("/movie/:id", handlers.DeleteMovieByIDHandler)
	r.PUT("/movie", handlers.WatchedMovieHandler)

	// listen and serve on 0.0.0.0:8080 and panics when error occurs.
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

// CORSMiddleware add middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
