package main

import (
	"fmt"
	"log"
	"mux-gorm/controllers"
	"net/http"

	"github.com/gorilla/mux"
	//menggunakan gorm versi baru
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/pahlawan", controllers.C_GetAllSuperheroBooks).Methods("GET")
	r.HandleFunc("/pahlawan/{name}", controllers.C_GetSuperheroById).Methods("GET")
	r.HandleFunc("/pahlawan", controllers.C_InsertSuperhero).Methods("POST")

	r.HandleFunc("/pahlawan/{name}", controllers.C_DeleteSuperheroById).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
