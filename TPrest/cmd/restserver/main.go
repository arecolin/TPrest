package main

import (
	"fmt"
	. "internal/web/rest"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	// Initialisation Database connection
	data.DbOpen("mydata.db")
	var path, err := data.DbPath()
	if err != nil {
		fmt.Println("erreur " + err)
	}
	data.DbClose()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/student", createStudent).Methods("POST")
	router.HandleFunc("/students", getAllStudents).Methods("GET")
	router.HandleFunc("/students/{id}", getOneStudent).Methods("GET")
	router.HandleFunc("/students/{id}", updateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")

	router.HandleFunc("/language", createLanguage).Methods("POST")
	router.HandleFunc("/languages", getAllLanguages).Methods("GET")
	router.HandleFunc("/language/{code}", getOneLanguage).Methods("GET")
	router.HandleFunc("/languages/{code}", updateLanguage).Methods("PUT")
	router.HandleFunc("/languages/{code}", deleteLanguage).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8080", router))
}