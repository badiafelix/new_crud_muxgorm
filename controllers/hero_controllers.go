package controllers

import (
	"encoding/json"
	"fmt"

	//"fmt"
	"mux-gorm/models"
	"net/http"

	//"strconv"
	"github.com/gorilla/mux"
	//menggunakan gorm versi baru
)

func C_GetAllSuperheroBooks(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func C_GetSuperheroById(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func C_InsertSuperhero(w http.ResponseWriter, r *http.Request) {
	var InsertHero models.Superhero
	var output models.OutputById
	_ = json.NewDecoder(r.Body).Decode(&InsertHero)
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func C_DeleteSuperheroById(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
