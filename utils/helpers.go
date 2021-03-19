package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateID(id string) bool  {
	return primitive.IsValidObjectID(id)
}
/* 
func IsValidAndExistID (w http.ResponseWriter, r *http.Request) bool {
	var response model.Response
	var logResponse model.LogResponse
	vars := mux.Vars(r)
	taskID := vars["id"]
	isvalidate := ValidateID(taskID)
	if !isvalidate {
		response.Message = model.INVALID_ID
		w.WriteHeader(http.StatusBadRequest)
		logResponse.Line = "123"
		logResponse.Message = model.INVALID_ID
		log.LogError(logResponse)
		json.NewEncoder(w).Encode(response)
		return false
	}
	err := per.ExistID(taskID)
	if err != nil {
		response.Message = model.INVALID_ID_NOT_EXIST
		w.WriteHeader(http.StatusBadRequest)
		logResponse.Line = "127"
		logResponse.Message = model.INVALID_ID_NOT_EXIST
		log.LogError(logResponse)
		json.NewEncoder(w).Encode(response)
		return false
	}
	return
} */