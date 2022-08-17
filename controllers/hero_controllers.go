package controllers

import (
	"encoding/json"
	"fmt"

	//"fmt"
	"mux-gorm/models"
	"net/http"

	"github.com/go-playground/validator/v10"

	//"strconv"
	"github.com/gorilla/mux"
	//menggunakan gorm versi baru
)

var validate *validator.Validate

func C_GetAllSuperheroBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var output models.OutputArray
	get_value, isValid, errorMessage := models.GetAllSuperhero()
	if isValid != true {
		output.Status = "error"
		output.Message = errorMessage.Error.Error()
	} else {
		output.Status = "success"
		output.Message = ""
		output.Data = get_value
	}
	json.NewEncoder(w).Encode(output)
}

func C_GetSuperheroById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	paramId := vars["name"]
	var output models.OutputById
	get_value, isValid, errorMessage := models.GetSuperheroById(paramId)
	if isValid != true {
		output.Status = "error"
		output.Message = errorMessage.Error.Error()
		output.Data = *get_value
	} else {
		output.Status = "success"
		output.Message = ""
		output.Data = *get_value
	}
	json.NewEncoder(w).Encode(output)
}

func C_InsertSuperhero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var InsertHero models.Superhero
	var output models.OutputById
	var outputError ErrorOutput
	_ = json.NewDecoder(r.Body).Decode(&InsertHero)
	//penempatan validator setelah decode
	validate = validator.New()
	err := validate.Struct(InsertHero)
	if err != nil {
		outputError.Status = "error"
		outputError.Message = err.Error()
		json.NewEncoder(w).Encode(outputError) //menampilkan message error input
	} else {
		_, isValid, errorMessage := models.InsertSuperhero(InsertHero)
		if isValid != true {
			output.Status = "error"
			output.Message = errorMessage.Error.Error()
		} else {
			checkData, isValid, errorMessage := models.GetSuperheroById(InsertHero.Name)
			if isValid != true {
				fmt.Printf(errorMessage.Error.Error())
			} else {
				output.Status = "success"
				output.Message = ""
				output.Data = *checkData
			}

		}
		json.NewEncoder(w).Encode(output)
	}

}

func C_DeleteSuperheroById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	paramId := vars["name"]
	var output models.OutputById
	get_value, isValid, errorMessage := models.DeleteSuperheroById(paramId)
	if isValid != true {
		output.Status = "error"
		output.Message = errorMessage.Error.Error()
		output.Data = *get_value
	} else {
		output.Status = "success"
		output.Message = "Data has been deleted"
		output.Data = *get_value
	}
	json.NewEncoder(w).Encode(output)
}

type ErrorOutput struct {
	Status  string
	Message string
}
