package movie

import (
	"context"
	"log"
	"movie-app/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Movie data structure
// each Movie struct will contain an ID, a Title, a Year and a Watched status.
// the second part of the struct shows how these 'fields' can be accessed in JSON.
type Movie struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Title   string             `json:"title"`
	Year    int                `json:"year"`
	Watched bool               `json:"watched"`
}

// GetAllMovies retrieves all movies from the db
func GetAllMovies() ([]*Movie, error) {
	var movies []*Movie

	collection := db.Client.Database("movies").Collection("movies")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	// This method will close the cursor after retrieving all documents.
	err = cursor.All(context.TODO(), &movies)
	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}
	return movies, nil
}

// GetMovieByID retrieves a movie by its id from the db
func GetMovieByID(id primitive.ObjectID) (*Movie, error) {
	var movie *Movie
	collection := db.Client.Database("movies").Collection("movies")
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&movie)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	return movie, err

}

// AddMovie adds a movie to the db
func AddMovie(movie *Movie) (primitive.ObjectID, error) {
	movie.ID = primitive.NewObjectID()
	result, err := db.Client.Database("movies").Collection("movies").InsertOne(context.TODO(), movie)
	if err != nil {
		log.Printf("Could not create movie: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

// DeleteMovieByID deletes a movie by its id from the db
func DeleteMovieByID(id primitive.ObjectID) error {
	collection := db.Client.Database("movies").Collection("movies")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Printf("Error: %v", err)
	}
	return nil
}

// WatchedMovie will put a movie its attribue 'watched' on true
func WatchedMovie(id primitive.ObjectID) error {
	collection := db.Client.Database("movies").Collection("movies")
	_, err := collection.UpdateOne(context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"watched", true}}},
		})
	if err != nil {
		log.Printf("Error: %v", err)
	}
	return nil
}
