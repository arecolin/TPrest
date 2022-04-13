package main

import (
	"TPrest/internal/entities"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/student", createStudent).Methods("POST")
	router.HandleFunc("/students", getAllStudents).Methods("GET")
	router.HandleFunc("/students/{id}", getOneStudent).Methods("GET")
	router.HandleFunc("/students/{id}", updateStudent).Methods("PATCH")
	router.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	var newEvent entities.Student
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	entities.Students = append(entities.Students, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func getOneStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	eventID, _ := strconv.ParseInt(id, 10, 0)

	for _, singleEvent := range entities.Students {
		if singleEvent.Id == int(eventID) {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllStudents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(entities.Students)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	studentId, _ := strconv.ParseInt(id, 10, 0)
	var updatedStudent entities.Student

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedStudent)

	for i, singleStudent := range entities.Students {
		if singleStudent.Id == int(studentId) {
			singleStudent.FirstName = updatedStudent.FirstName
			singleStudent.LastName = updatedStudent.LastName
			singleStudent.Age = updatedStudent.Age
			singleStudent.LanguageCode = updatedStudent.LanguageCode
			entities.Students = append(entities.Students[:i], singleStudent)
			json.NewEncoder(w).Encode(singleStudent)
		}
	}
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	StudentId, _ := strconv.ParseInt(id, 10, 0)

	for i, singleStudent := range entities.Students {
		if singleStudent.Id == int(StudentId) {
			entities.Students = append(entities.Students[:i], entities.Students[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", StudentId)
		}
	}
}
