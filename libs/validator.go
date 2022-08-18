package libs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mux-gorm/models"

	"os"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10" //validator menggunakan go-playground/validator/v10
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

type ErrorOutput struct {
	Status  string
	Message string
}

func CobaValidator(params models.Superhero) string {
	var errorMsg string
	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validate = validator.New()
	InputValidation(trans)
	err := validate.Struct(params)
	if err != nil {

		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			// can translate each error one at a time.
			fmt.Println(e.Translate(trans))
			errorMsg = e.Translate(trans)
		}
	}

	return errorMsg
}

func InputValidation(trans ut.Translator) {

	// Membuka file json
	jsonFile, err := os.Open("app-setting/message.json")
	// buka os.Open returns an error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened message.json")
	// defer closing json File
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	hasil := result["errorMessage"].([]interface{})
	hasil2 := result["warningMessage"].(interface{})
	//fmt.Println(hasil[0].(map[string]interface{})["max"])
	//fmt.Println(hasil[0].(map[string]interface{})["min"])
	//fmt.Println(hasil[0].(map[string]interface{})["required"])
	fmt.Println(hasil2.(map[string]interface{})["maju"])

	required := hasil[0].(map[string]interface{})["required"].(string)
	min := hasil[0].(map[string]interface{})["min"].(string)
	max := hasil[0].(map[string]interface{})["max"].(string)
	//fmt.Println(result["errorMessage"])
	////////////////////////////////
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", required, true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", min, true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field())

		return t
	})

	validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", max, true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field())

		return t
	})
}
