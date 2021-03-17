package services

import (
	"../log"
	"../model"
	per "../persistence"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const FILE_NAME = "service.go"

var logResponse = model.LogResponse{FILE_NAME, "", ""}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response model.Response
	var newTask model.Task
	res, err := ioutil.ReadAll(r.Body)
	if err != nil || len(res) == 0 {
		response.Message = model.INVALID_BODY
		w.WriteHeader(http.StatusBadRequest)
		logResponse.Line = "20"
		logResponse.Message = model.INVALID_BODY
		log.LogError(logResponse)
		json.NewEncoder(w).Encode(response)
		return
	}
	json.Unmarshal(res, &newTask)
	newTask.CreatedAt = time.Now()
	newTask.UpdatedAt = time.Now()
	err = per.CreateTask(&newTask)
	if err != nil {
		response.Message = model.SAVE_NEW_DATA_ERROR
		response.Payload = &newTask
		logResponse.Message = model.SAVE_NEW_DATA_ERROR
		logResponse.Line = "33"
		log.LogError(logResponse)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Message = model.SAVE_NEW_DATA_SUCCESS
	response.Payload = &newTask
	json.NewEncoder(w).Encode(response)
}

func GetAll(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response model.Response
	var tasks, err = per.GetAll()
	if err != nil {
		response.Message = model.GET_ALL_ERROR
		logResponse.Message = model.GET_ALL_ERROR
		logResponse.Line = "49"
		log.LogError(logResponse)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Message = model.GET_ALL_SUCCESS
	response.PayloadList = tasks
	json.NewEncoder(w).Encode(response)
}
