package main

import (
	"fmt"
	data "internal/bdd"
	entities "internal/entities"
	. "internal/web/rest"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Initialisation Database connection
	data.DbOpen("mydata.db")

	// Database PATH
	var path = data.DbPath()
	fmt.Println("Database path: " + path)

	// Create Bucket
	data.CreateBucket("languages")
	data.CreateBucket("students")

	// Playground for testing
	// Create a new language / student and save it to the database
	var student = entities.Student{Id: "1", FirstName: "John", LastName: "Doe", Age: 30, LanguageCode: "en"}
	data.SaveStudent(student)
	var language = entities.Language{Code: "en", Name: "English"}
	data.SaveLanguage(language)
	// Get student / language from database
	var getStudent = data.DbGetStudent("1")
	fmt.Println("Student: " + getStudent)
	var getLanguage = data.DbGetLanguage("en")
	fmt.Println("Language: " + getLanguage)
	// Update student / language in database
	var updateStudent = entities.Student{Id: "1", FirstName: "John", LastName: "Doe", Age: 25, LanguageCode: "fr"}
	data.DbUpdateStudent(updateStudent)
	var updateLanguage = entities.Language{Code: "en", Name: "Anglais"}
	data.DbUpdateLanguage(updateLanguage)
	// GetAll students / languages from database
	var getAllStudents = data.DbGetAll("students")
	fmt.Println("All students: ", getAllStudents)
	var getAllLanguages = data.DbGetAll("languages")
	fmt.Println("All languages: ", getAllLanguages)
	// Delete student / language from database
	data.DbDeleteStudent("1")
	fmt.Println(data.DbGetAll("students"))
	data.DbDeleteLanguage("en")
	fmt.Println(data.DbGetAll("languages"))

	// Initialisation Web Server
	fmt.Println("Starting web server at :8080")
	router := mux.NewRouter().StrictSlash(true)

	// Routes for Students
	fmt.Println("Routes for Students...")
	router.HandleFunc("/student", CreateStudent).Methods("POST")
	router.HandleFunc("/students", GetAllStudents).Methods("GET")
	router.HandleFunc("/student/{id}", GetOneStudent).Methods("GET")
	router.HandleFunc("/students/{id}", UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", DeleteStudent).Methods("DELETE")

	// Routes for Languages
	fmt.Println("Routes for Languages...")
	router.HandleFunc("/language", CreateLanguage).Methods("POST")
	router.HandleFunc("/languages", GetAllLanguages).Methods("GET")
	router.HandleFunc("/language/{code}", GetOneLanguage).Methods("GET")
	router.HandleFunc("/languages/{code}", UpdateLanguage).Methods("PUT")
	router.HandleFunc("/languages/{code}", DeleteLanguage).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
