package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm" //menggunakan gorm versi baru
)

type superhero struct { //automigrate akan membuat nama tabel di db dengan nama seperti nama struct
	gorm.Model
	Name   string
	Value  string
	Author string
}

var db *gorm.DB

func connectDb() {
	dsn := "host=localhost user=postgres password=root1234 dbname=Go_database port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	db = conn
}

func getAllSuperhero(w http.ResponseWriter, r *http.Request) {
	connectDb()
	var sel []superhero
	db.Raw("SELECT * FROM test_automigrates").Scan(&sel)
	fmt.Println(sel)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sel)
}

func getSuperheroById(w http.ResponseWriter, r *http.Request) {
	connectDb()
	params := mux.Vars(r)
	var selById superhero
	db.Raw("SELECT * FROM test_automigrates WHERE name = ?", params["name"]).Scan(&selById)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(selById)
}

func insertSuperhero(w http.ResponseWriter, r *http.Request) {
	connectDb()
	var InsertHero superhero
	_ = json.NewDecoder(r.Body).Decode(&InsertHero)
	db.Raw("INSERT INTO test_automigrates(name,value,author)VALUES(?, ?, ?);", InsertHero.Name, InsertHero.Value, InsertHero.Author).Scan(&InsertHero)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(InsertHero)

}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/pahlawan", getAllSuperhero).Methods("GET")
	r.HandleFunc("/pahlawan/{name}", getSuperheroById).Methods("GET")
	r.HandleFunc("/pahlawan", insertSuperhero).Methods("POST")
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

	// //update table menggunakan raw sql
	// var updt test_automigrate
	// db.Raw("UPDATE test_automigrates SET value = ? WHERE name = ? RETURNING name, value, author", "putri", "sri asih").Scan(&updt)
	// fmt.Println(updt)
}
