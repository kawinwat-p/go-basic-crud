package repositories

import (
	"context"
	"errors"
	"fmt"
	"os"
	. "practice/domain/datasources"
	"practice/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type artistsRepository struct {
	Context context.Context
	Collection *mongo.Collection
}

type IArtistsRepository interface {
	GetAllArtists() ([]entities.ArtistDataFormat, error)
	GetArtistByName(name string) (*entities.ArtistDataFormat, error)
	// ArtistExist(name string) (bool, error)
	// UpdateArtistByName(name string, data entities.ArtistDataFormat) error
	// CreateArtist(data entities.ArtistDataFormat) error
	// DeleteArtistByName(name string) error
}

func NewArtistsRepository(db *MongoDB) IArtistsRepository {
	return &artistsRepository{
		Context:    db.Context,
		Collection: db.MongoDB.Database(os.Getenv("DATABASE_NAME")).Collection("artists"),
	}
}

func (repo artistsRepository) GetAllArtists() ([]entities.ArtistDataFormat, error){
	result := []entities.ArtistDataFormat{}

	cursor ,err := repo.Collection.Find(repo.Context, bson.M{},options.Find())

	if err != nil {
		return nil,err
	}
	defer cursor.Close(repo.Context)

	for cursor.Next(repo.Context) {
		var artist entities.ArtistDataFormat

		err = cursor.Decode(&artist)
		if err != nil {
			fmt.Println("cannot get artist repo")
			return nil,err
		}
		result = append(result,artist)
	}

	return result,nil
}

func (repo artistsRepository) GetArtistByName(name string) (*entities.ArtistDataFormat, error) {
	result := entities.ArtistDataFormat{}
	filter := bson.M{"artist_name": name}

	data := repo.Collection.FindOne(repo.Context,filter).Decode(&result)

	if data == mongo.ErrNoDocuments{
		return nil,errors.New("artist not found")
	}

	return &result, nil
}