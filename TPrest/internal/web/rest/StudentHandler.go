package web

import (
	"encoding/json"
	"fmt"
	data "internal/bdd"
	. "internal/entities"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &student)
	Students = append(Students, student)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(student)
	data.SaveStudent(student)
}

func GetOneStudent(w http.ResponseWriter, r *http.Request) {
	studentID := mux.Vars(r)["id"]
	for _, singleStudent := range Students {
		if singleStudent.Id == studentID {
			json.NewEncoder(w).Encode(singleStudent)
			data.DbGetStudent(studentID)
		}
	}
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Students)
	data.DbGetAll("students")
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	studentId := id
	var updatedStudent Student

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedStudent)

	for i, singleStudent := range Students {
		if singleStudent.Id == studentId {
			singleStudent.FirstName = updatedStudent.FirstName
			singleStudent.LastName = updatedStudent.LastName
			singleStudent.Age = updatedStudent.Age
			singleStudent.LanguageCode = updatedStudent.LanguageCode
			Students = append(Students[:i], singleStudent)
			json.NewEncoder(w).Encode(singleStudent)

			data.DbUpdateStudent(singleStudent)
		}
	}
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	StudentId := mux.Vars(r)["id"]

	for i, singleStudent := range Students {
		if singleStudent.Id == StudentId {
			Students = append(Students[:i], Students[i+1:]...)

			data.DbDeleteStudent(singleStudent.Id)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", StudentId)
		}
	}
}
