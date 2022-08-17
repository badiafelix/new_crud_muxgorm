package controllers

import (
	"encoding/json"
	"fmt"

	//"fmt"
	"mux-gorm/models"
	"net/http"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	//"strconv"
	"github.com/gorilla/mux"
	//menggunakan gorm versi baru
)

//var validate *validator.Validate
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
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

func InputValidation(trans ut.Translator) {
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} harus diisi!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "{0} Kurang banyak oi!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field())

		return t
	})

	validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "{0} Kebanyakan oi!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field())

		return t
	})
}

func C_InsertSuperhero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var InsertHero models.Superhero
	var output models.OutputById
	var outputError ErrorOutput
	_ = json.NewDecoder(r.Body).Decode(&InsertHero)
	//penempatan validator setelah decode
	//validate = validator.New()

	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validate = validator.New()
	InputValidation(trans)
	err := validate.Struct(InsertHero)
	if err != nil {
		var tmp string
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			// can translate each error one at a time.
			fmt.Println(e.Translate(trans))
			tmp = e.Translate(trans)
		}
		outputError.Status = "error"
		outputError.Message = tmp
		//outputError.Message = err.Error()
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
