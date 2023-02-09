package api

import (
	"Microservice/logger/models"
	"Microservice/logger/utils"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func WriteLogs(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var log models.Log
		err := json.NewDecoder(r.Body).Decode(&log)
		utils.HandleErrors(err)

		_, err = log.Insert(db)
		utils.HandleErrors(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Log inserted successfully"))

	}
}
