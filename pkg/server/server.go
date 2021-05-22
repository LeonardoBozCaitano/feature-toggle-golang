package server

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Router     *mux.Router
	Collection *mongo.Collection
}

func NewServer() (*Server, error) {
	database, err := ConnectToDatabase()

	if err != nil {
		return nil, err
	}
	server := &Server{
		Router: mux.NewRouter(), Collection: database.Collection("features"),
	}
	server.routes()
	return server, nil
}
