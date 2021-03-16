package main

import (
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

//MONGO_DB_ATLAS=mongodb+srv://{user}:{password}@clustertest.loyxx.mongodb.net/{name_bd}
func indexRoute(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Welcome to my api")
}
func getTasks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		fmt.Fprintf(w, "Ingrese datos correctos")

	}
	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

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

	reqBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
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
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", deletTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))
}
