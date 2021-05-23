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

func TestFeatureGetAllShouldGetAllThree(t *testing.T) {
	service := NewTestService()

	service.Insert(&feature.FeatureEntity{primitive.NewObjectID(), "feature1", []string{"client1", "client2"}})
	service.Insert(&feature.FeatureEntity{primitive.NewObjectID(), "feature2", []string{"client1"}})
	service.Insert(&feature.FeatureEntity{primitive.NewObjectID(), "feature3", []string{"client1"}})

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

func TestFeatureInsertShouldReturnNameAlreadyExistsError(t *testing.T) {
	service := NewTestService()
	service.Insert(&feature.FeatureEntity{primitive.NewObjectID(), "feature1", []string{"client1", "client2"}})

	input := &feature.FeatureEntity{primitive.NewObjectID(), "feature1", []string{"client1"}}

	_, err := service.Insert(input)

	if err != nil && err.Error() != "Feature name already exists" {
		t.Error("Expected error not found.")
	} else {
		t.Logf("Success!!")
	}

	service.GetFeatureCollection().DeleteMany(nil, bson.D{})
}

func TestFeatureInsertShouldInsertSuccessfully(t *testing.T) {
	service := NewTestService()

	input := &feature.FeatureEntity{primitive.NewObjectID(), "feature1", []string{"client1"}}

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

func TestValidateFeatureClientShouldReturnEnabledTrue(t *testing.T) {
	service := NewTestService()

	inputFeatureName := "feature1"
	inputClientName := "client1"
	service.Insert(&feature.FeatureEntity{primitive.NewObjectID(), inputFeatureName, []string{inputClientName}})

	output, err := service.ValidateFeatureClient(inputFeatureName, inputClientName)

	if err != nil {
		t.Errorf("Error: %v", err)
	} else {
		if cmp.Equal(true, *output) {
			t.Logf("Success!!")
		} else {
			t.Errorf("Error on output: %v", *output)
		}
	}

	service.GetFeatureCollection().DeleteMany(nil, bson.D{})
}

func TestValidateFeatureClientShouldReturnEnabledFalse(t *testing.T) {
	service := NewTestService()

	inputFeatureName := "feature1"
	inputClientName := "client1"
	service.Insert(&feature.FeatureEntity{primitive.NewObjectID(), inputFeatureName, []string{"client2", "client3"}})

	output, err := service.ValidateFeatureClient(inputFeatureName, inputClientName)

	if err != nil {
		t.Errorf("Error: %v", err)
	} else {
		if cmp.Equal(false, *output) {
			t.Logf("Success!!")
		} else {
			t.Errorf("Error on output: %v", *output)
		}
	}

	service.GetFeatureCollection().DeleteMany(nil, bson.D{})
}
