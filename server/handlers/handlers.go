package handlers

import (
	"log"
	"movie-app/movie"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllMoviesHandler to get all movies
func GetAllMoviesHandler(c *gin.Context) {
	loadedMovies, err := movie.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, loadedMovies) //gin.H{loadedMovies})
}

// GetMovieByIDHandler to get a movie by ID
// func GetMovieByIDHandler(c *gin.Context) {
// 	movieID, err := primitive.ObjectIDFromHex(c.Param("id"))
// 	if err != nil {
// 		log.Println("Invalid id")
// 	}
// 	loadedMovie, err := movie.GetMovieByID(movieID)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"msg": err})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"id": loadedMovie.ID, "title": loadedMovie.Title,
// 		"year": loadedMovie.Year, "watched": loadedMovie.Watched})
// }

// AddMovieHandler to add a movie
func AddMovieHandler(c *gin.Context) {
	var mov movie.Movie
	if err := c.ShouldBindJSON(&mov); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := movie.AddMovie(&mov)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// DeleteMovieByIDHandler to get a movie by ID
func DeleteMovieByIDHandler(c *gin.Context) {
	movieID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		log.Println("Invalid id")
	}
	if err := movie.DeleteMovieByID(movieID); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "")
}

// WatchedMovieHandler to put a movie on watched
func WatchedMovieHandler(c *gin.Context) {
	var mov movie.Movie
	if err := c.ShouldBindJSON(&mov); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	err := movie.WatchedMovie(mov.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, "")
}
