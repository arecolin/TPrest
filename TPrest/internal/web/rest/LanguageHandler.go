package web

import (
	
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func CreateLanguage(w http.ResponseWriter, r *http.Request) {
	var language Language
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &language)
	Languages = append(Languages, language)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(language)
}

func GetOneLanguage(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["Code"]

	languageCode := code

	for _, singleLanguage := range Languages {
		if singleLanguage.Code == languageCode {
			json.NewEncoder(w).Encode(singleLanguage)
		}
	}
}

func GetAllLanguages(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Languages)
}

func UpdateLanguage(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["Code"]

	languageCode := code
	var updatedLanguage Language

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedLanguage)

	for i, singleLanguage := range Languages {
		if singleLanguage.Code == languageCode {
			singleLanguage.Name = updatedLanguage.Name
			Languages = append(Languages[:i], singleLanguage)
			json.NewEncoder(w).Encode(singleLanguage)
		}
	}
}

func DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["Code"]

	languageCode := code

	for i, singleLanguage := range Languages {
		if singleLanguage.Code == languageCode {
			Languages = append(Languages[:i], Languages[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", languageCode)
		}
	}
}