package services

import (
	"fmt"
	"practice/domain/entities"
	"practice/domain/repositories"
)

type artistService struct {
	ArtistRepository repositories.IArtistsRepository
}

type IArtistService interface {
	GetAllArtistsService() ([]entities.ArtistDataFormat, error)
	GetArtistByNameService(name string) (*entities.ArtistDataFormat, error)
}

func NewArtistService(artistRepository repositories.IArtistsRepository) IArtistService {
	return &artistService{
		ArtistRepository: artistRepository,
	}
}

func (sv artistService) GetAllArtistsService() ([]entities.ArtistDataFormat, error) {

	result, err := sv.ArtistRepository.GetAllArtists()

	if err != nil {
		fmt.Println("cannot get artist service")
		return nil, err
	}

	return result, nil
}

func (sv artistService) GetArtistByNameService(name string) (*entities.ArtistDataFormat, error) {
	result, err := sv.ArtistRepository.GetArtistByName(name)

	if err != nil {
		return nil, err
	}

	return result,nil
}