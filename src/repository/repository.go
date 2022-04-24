package repository

import (
	"github.com/madrid-recicla/server/src/model"
	"go.mongodb.org/mongo-driver/bson"
)

type ContainerRepository interface {
	ListAll() (*[]*model.Location, error)
	ListAllMatching(filter *bson.D) (*[]*model.Location, error)
	AddAll(containers *[]*model.Location) (int, error)
	DeleteAll() (int, error)
	DeleteAllMatching(filter *bson.D) (int, error)
}
