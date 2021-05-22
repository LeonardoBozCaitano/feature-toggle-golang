package feature_test

import (
	"testing"

	"github.com/feature_toggle/pkg/feature"
	"github.com/feature_toggle/pkg/server"
	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewTestService() *feature.Service {
	database, _ := server.ConnectToDatabase()
	return feature.NewService(database.Collection("feature_test"))
}

func TestFeatureGetAll(t *testing.T) {
	service := NewTestService()

	service.Insert(&feature.FeatureEntity{primitive.NewObjectID(), "feature1", []int{1, 2, 3}})
	service.Insert(&feature.FeatureEntity{primitive.NewObjectID(), "feature2", []int{2}})
	service.Insert(&feature.FeatureEntity{primitive.NewObjectID(), "feature3", []int{2}})

	output, err := service.GetAll()
	t.Logf("Success!!")

	if err != nil {
		t.Errorf("error found: %v", err)
	}

	if len(output) == 3 {
		t.Logf("Success!!")
	} else {
		t.Errorf("wrong features found: %v", len(output))
	}

	service.GetFeatureCollection().DeleteMany(nil, bson.D{})
}

func TestFeatureInsert(t *testing.T) {
	service := NewTestService()

	input := &feature.FeatureEntity{primitive.NilObjectID, "feature1", []int{1, 2, 3}}

	output, err := service.Insert(input)

	if err != nil {
		t.Errorf("error found: %v", err)
	}

	validate := service.GetFeatureCollection().FindOne(nil, bson.D{primitive.E{Key: "name", Value: input.Name}})

	var validationFeature *feature.FeatureEntity
	err = validate.Decode(&validationFeature)

	if cmp.Equal(validationFeature, output) {
		t.Logf("Success!!")
	} else {
		t.Errorf("wrong features found: %v != %v", validationFeature, output)
	}

	service.GetFeatureCollection().DeleteMany(nil, bson.D{})
}
