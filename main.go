package main

import (
	service "./services"
	ut "./utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type task struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type (
	allTask []task
)

var tasks = allTask{
	{
		ID:      1,
		Name:    "Task 1",
		Content: "content 1",
	},
}


func deletTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Id invalido")
		return
	}
	var taskResponse task
	for index, task := range tasks {
		if task.ID == taskID {
			taskResponse = task
			tasks = append(tasks[:index], tasks[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(taskResponse)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Id no existe")
}


func main() {

	router := mux.NewRouter().StrictSlash(true)
	ut.EnableCORS(router)
	router.HandleFunc("/tasks", service.GetAll).Methods("GET")
	router.HandleFunc("/tasks", service.Create).Methods("POST")
	router.HandleFunc("/tasks/{id}", ut.IsValidAndExistID(service.Delete)).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", ut.IsValidAndExistID(service.Completed)).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))

}
