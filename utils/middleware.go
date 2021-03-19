package utils

import (
	"encoding/json"
	"net/http"

	"../log"
	"../model"
	per "../persistence"
	"github.com/gorilla/mux"
)

func EnableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
 		    w.Header().Set("Access-Control-Allow-Origin", "http://localhost")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			next.ServeHTTP(w, req)
		})
}


func IsValidAndExistID(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var response model.Response
		var logResponse model.LogResponse
		vars := mux.Vars(req)
		taskID := vars["id"]
		isvalidate := ValidateID(taskID)
		if !isvalidate {
			response.Message = model.INVALID_ID
			res.WriteHeader(http.StatusBadRequest)
			logResponse.Line = "123"
			logResponse.Message = model.INVALID_ID
			log.LogError(logResponse)
			json.NewEncoder(res).Encode(response)
			
		}
		err := per.ExistID(taskID)
		if err != nil {
			response.Message = model.INVALID_ID_NOT_EXIST
			res.WriteHeader(http.StatusBadRequest)
			logResponse.Line = "127"
			logResponse.Message = model.INVALID_ID_NOT_EXIST
			log.LogError(logResponse)
			json.NewEncoder(res).Encode(response)
		}
	  next(res, req)
	}
  }