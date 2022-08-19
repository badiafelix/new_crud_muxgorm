package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mux-gorm/config"
	"net/http"

	"gorm.io/gorm" //menggunakan gorm versi baru
)

type Superhero struct { //automigrate akan membuat nama tabel di db dengan nama seperti nama struct
	gorm.Model
	Name   string `validate:"required,min=3,max=15"`
	Value  string `validate:"required,min=3,max=15"`
	Author string
}

type OutputById struct {
	Status  string
	Message string
	Data    Superhero
}

type OutputArray struct {
	Status  string
	Message string
	Data    []Superhero
}

func GetAllSuperhero() ([]Superhero, bool, *gorm.DB) {
	var Sel []Superhero
	var isValid bool
	config.ConnectDb()
	err := config.Db.Raw("SELECT * FROM test_automigrates").Scan(&Sel)
	if err.Error != nil {
		isValid = false
		fmt.Println("ada error")
		fmt.Printf("pesan errornya adalah : %v", err.Error.Error())
	} else {
		isValid = true
		fmt.Println("tidak ada error")
	}
	return Sel, isValid, err
}

func GetSuperheroById(paramName string) (*Superhero, bool, *gorm.DB) {
	var selById Superhero
	var isValid bool
	config.ConnectDb()
	err := config.Db.Raw("SELECT * FROM test_automigrates WHERE name = ?", paramName).Scan(&selById)
	if err.Error != nil {
		isValid = false
		fmt.Println("ada error")
		fmt.Printf("pesan errornya adalah : %v", err.Error.Error())
	} else {
		isValid = true
		fmt.Println("tidak ada error")
	}
	return &selById, isValid, err
}

func InsertSuperhero(paramINput Superhero) (*Superhero, bool, *gorm.DB) { // return 3 type data *superhero, isValid dan *gorm.db
	var InsertHero Superhero
	var isValid bool
	config.ConnectDb()
	err := config.Db.Raw("INSERT INTO test_automigrates(created_at,name,value,author)VALUES(current_timestamp,?, ?, ?);", paramINput.Name, paramINput.Value, paramINput.Author).Scan(&InsertHero)
	if err.Error != nil {
		isValid = false
		fmt.Println("ada error")
		fmt.Printf("pesan errornya adalah : %v", err.Error.Error())
	} else {
		isValid = true
		fmt.Println("tidak ada error")
		//GetSuperheroById(paramINput.Name)
	}
	return &InsertHero, isValid, err
}

func UpdateSuperhero(paramINput Superhero) (*Superhero, bool, *gorm.DB) { // return 3 type data *superhero, isValid dan *gorm.db
	var InsertHero Superhero
	var isValid bool
	config.ConnectDb()
	//err := config.Db.Raw("INSERT INTO test_automigrates(created_at,name,value,author)VALUES(current_timestamp,?, ?, ?);", paramINput.Name, paramINput.Value, paramINput.Author).Scan(&InsertHero)
	err := config.Db.Raw("UPDATE test_automigrates SET updated_at = current_timestamp, value = ?, author = ? WHERE name = ?;", paramINput.Value, paramINput.Author, paramINput.Name).Scan(&InsertHero)

	if err.Error != nil {
		isValid = false
		fmt.Println("ada error")
		fmt.Printf("pesan errornya adalah : %v", err.Error.Error())
	} else {
		isValid = true
		fmt.Println("tidak ada error")
		//GetSuperheroById(paramINput.Name)
	}
	return &InsertHero, isValid, err
}

func DeleteSuperheroById(paramName string) (*Superhero, bool, *gorm.DB) {
	var selById Superhero
	var isValid bool
	config.ConnectDb()
	err := config.Db.Raw("DELETE FROM test_automigrates WHERE name = ?", paramName).Scan(&selById)
	if err.Error != nil {
		isValid = false
		fmt.Println("ada error")
		fmt.Printf("pesan errornya adalah : %v", err.Error.Error())
	} else {
		isValid = true
		fmt.Println("tidak ada error")
	}
	return &selById, isValid, err
}

// func InsertSuperhero(w http.ResponseWriter, r *http.Request) {
// 	config.ConnectDb()
// 	var InsertHero Superhero
// 	_ = json.NewDecoder(r.Body).Decode(&InsertHero)
// 	config.Db.Raw("INSERT INTO test_automigrates(name,value,author)VALUES(?, ?, ?);", InsertHero.Name, InsertHero.Value, InsertHero.Author).Scan(&InsertHero)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(InsertHero)

// }

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
