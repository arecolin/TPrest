package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	. "entities"
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
}

func GetOneStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["Id"]

	studentID, _ := strconv.ParseInt(id, 10, 0)

	for _, singleStudent := range Students {
		if singleStudent.Id == int(studentID) {
			json.NewEncoder(w).Encode(singleStudent)
		}
	}
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Students)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["Id"]

	studentId, _ := strconv.ParseInt(id, 10, 0)
	var updatedStudent Student

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedStudent)

	for i, singleStudent := range Students {
		if singleStudent.Id == int(studentId) {
			singleStudent.FirstName = updatedStudent.FirstName
			singleStudent.LastName = updatedStudent.LastName
			singleStudent.Age = updatedStudent.Age
			singleStudent.LanguageCode = updatedStudent.LanguageCode
			Students = append(Students[:i], singleStudent)
			json.NewEncoder(w).Encode(singleStudent)
		}
	}
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["Id"]

	StudentId, _ := strconv.ParseInt(id, 10, 0)

	for i, singleStudent := range Students {
		if singleStudent.Id == int(StudentId) {
			Students = append(Students[:i], Students[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", StudentId)
		}
	}
}