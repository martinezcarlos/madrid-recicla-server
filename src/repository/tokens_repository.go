package repository

import (
	"log"
	"os"
)

type mapboxTokenRepository struct {
}

func NewMapboxTokenRepository() TokenRepository {
	return &mapboxTokenRepository{}
}

func (r *mapboxTokenRepository) GetToken() (string, error) {
	log.Println("Fetching MapBox token")
	// For now, return the token from the environment variable.
	return os.Getenv("MAPBOX_TOKEN"), nil
}
