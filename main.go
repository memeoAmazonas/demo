package main

import (
	service "./services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
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

func indexRoute(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Welcome to my api")
}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Id invalido")
		return
	}
	var taskResponse task
	for _, task := range tasks {
		if task.ID == taskID {
			taskResponse = task
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(taskResponse)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Id no existe")
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
func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Id invalido")
		return
	}

	reqBody, err2 := ioutil.ReadAll(r.Body)
	if err2 != nil {
		fmt.Fprintf(w, "Datos no validos")
	}
	var newTask task

	json.Unmarshal(reqBody, &newTask)

	for i, t := range tasks {
		if t.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			newTask.ID = taskID
			tasks = append(tasks, newTask)
			fmt.Fprintf(w, "La tarea con el id %v fue actualizada", taskID)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Id no existe")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", service.GetAll).Methods("GET")
	router.HandleFunc("/tasks", service.Create).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", deletTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))

}
