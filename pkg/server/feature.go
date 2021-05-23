package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/feature_toggle/pkg/feature"
	"github.com/gorilla/mux"
)

type featureResponse struct {
	Name    string   `json:"name"`
	Clients []string `json:"clients"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func (t *Server) HandleFeatureGetAll() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		service := feature.NewService(t.Collection)
		result, err := service.GetAll()

		if err != nil {
			log.Printf("Error while handling insert feature: %s", err)
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(&errorResponse{
				Message: err.Error(),
			})
			return
		}

		var featureResponseList []featureResponse

		for _, element := range result {
			t := featureResponse{
				Name:    element.Name,
				Clients: element.Clients,
			}
			featureResponseList = append(featureResponseList, t)
		}

		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(featureResponseList)
	}
}

func (t *Server) HandleFeatureInsert() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		input := &feature.FeatureEntity{}
		json.NewDecoder(req.Body).Decode(input)

		service := feature.NewService(t.Collection)
		result, err := service.Insert(input)

		if err != nil {
			log.Printf("Error while handling insert feature: %s", err)
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(&errorResponse{
				Message: err.Error(),
			})
			return
		}

		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(&featureResponse{
			Name:    result.Name,
			Clients: result.Clients,
		})
	}
}

func (t *Server) HandleFeatureClientVerification() http.HandlerFunc {
	type response struct {
		Enabled bool `json:"enable"`
	}
	return func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		inputFeatureName := vars["name"]
		inputClientName := vars["client"]

		service := feature.NewService(t.Collection)
		result, err := service.ValidateFeatureClient(inputFeatureName, inputClientName)

		if err != nil {
			log.Printf("Error while handling insert feature: %s", err)
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(&errorResponse{
				Message: err.Error(),
			})
			return
		}

		res.WriteHeader(http.StatusFound)
		json.NewEncoder(res).Encode(&response{
			Enabled: *result,
		})
	}
}
