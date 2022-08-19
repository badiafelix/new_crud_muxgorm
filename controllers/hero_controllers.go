package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mux-gorm/libs"
	"mux-gorm/models"
	"net/http"
	"os"

	//"strconv"
	"github.com/gorilla/mux"
	//menggunakan gorm versi baru
)

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
	var outputError libs.ErrorOutput
	_ = json.NewDecoder(r.Body).Decode(&InsertHero)
	validasi := libs.CobaValidator(InsertHero)
	if len(validasi) > 0 {
		outputError.Status = "error"
		outputError.Message = validasi
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

	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

}

func C_UpdateSuperhero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var InsertHero models.Superhero
	var output models.OutputById
	var outputError libs.ErrorOutput
	_ = json.NewDecoder(r.Body).Decode(&InsertHero)
	validasi := libs.CobaValidator(InsertHero)
	if len(validasi) > 0 {
		outputError.Status = "error"
		outputError.Message = validasi
		json.NewEncoder(w).Encode(outputError) //menampilkan message error input
	} else {
		_, isValid, errorMessage := models.UpdateSuperhero(InsertHero)
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
