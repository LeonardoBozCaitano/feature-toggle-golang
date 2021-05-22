package feature

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FeatureEntity struct {
	Id      primitive.ObjectID `bson:"_id"`
	Name    string             `bson:"name"`
	Clients []int              `bson:"clients"`
}

type Service struct {
	collection *mongo.Collection
}

func NewService(collection *mongo.Collection) *Service {
	return &Service{
		collection: collection,
	}
}

func (service *Service) GetAll() ([]*FeatureEntity, error) {
	var features []*FeatureEntity

	cur, err := service.collection.Find(nil, bson.D{})
	if err != nil {
		return features, err
	}

	for cur.Next(nil) {
		var t FeatureEntity
		err := cur.Decode(&t)
		if err != nil {
			return features, err
		}

		features = append(features, &t)
	}

	if err := cur.Err(); err != nil {
		return features, err
	}

	cur.Close(nil)

	if len(features) == 0 {
		return features, mongo.ErrNoDocuments
	}

	return features, nil
}

func (service *Service) Insert(feature *FeatureEntity) (*FeatureEntity, error) {
	feature.Id = primitive.NewObjectID()
	_, err := service.collection.InsertOne(nil, feature)

	return feature, err
}

func (service *Service) GetFeatureCollection() (collection *mongo.Collection) {
	return service.collection
}
