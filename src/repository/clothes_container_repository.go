package repository

import (
	"log"

	"github.com/madrid-recicla/server/src/dbconnector"
	"github.com/madrid-recicla/server/src/model"
	"go.mongodb.org/mongo-driver/bson"
)

type clothesRepository struct {
}

func NewClothesRepository() ContainerRepository {
	return &clothesRepository{}
}

func (r *clothesRepository) ListAll() (*[]*model.Location, error) {
	return r.ListAllMatching(&bson.D{})
}

func (r *clothesRepository) ListAllMatching(filter *bson.D) (*[]*model.Location, error) {
	log.Println("Retrieving clothes containers for filter", *filter)
	documents, err := dbconnector.
		ListDocuments[model.Location](dbconnector.COLLECTION_CLOTHES_CONTAINER, filter)
	if err != nil {
		log.Println("Error while searching ClothContainers:", err)
		return nil, err
	}
	return documents, nil
}

func (r *clothesRepository) AddAll(containers *[]*model.Location) (int, error) {
	inserted, err := dbconnector.InsertDocuments(dbconnector.COLLECTION_CLOTHES_CONTAINER, containers)
	if err != nil {
		log.Println("Error while adding ClothContainers:", err)
		return 0, err
	}
	log.Printf("Inserted %d ClothContainers", len(inserted.InsertedIDs))
	return len(inserted.InsertedIDs), nil
}

func (r *clothesRepository) DeleteAll() (int, error) {
	return r.DeleteAllMatching(&bson.D{})
}

func (r *clothesRepository) DeleteAllMatching(filter *bson.D) (int, error) {
	log.Println("Deleting clothes containers for filter", *filter)
	deleted, err := dbconnector.DeleteDocuments(dbconnector.COLLECTION_CLOTHES_CONTAINER, filter)
	if err != nil {
		log.Println("Error while deleting ClothContainers:", err)
		return 0, err
	}
	return int(deleted.DeletedCount), nil
}
