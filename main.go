package main

import{
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
}

type movie struct{
	id string `json:  "id"`
	isbn string `json: "isbn"`
	title string `json: "title"`
	director *director `json: "director"`
}

type director struct{
	firstName string `json: "firstname"`
	lastName string `json: "lastname"`
}

var movies []movie

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",  deleteMovie).Methods("DELETE")
	
	fmt.Printf("Starting server at Port 8000")
	log.Fatal(http.ListenAndServer(":8000",r))
}	