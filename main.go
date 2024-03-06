package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

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

//The function below converts the incoming 
//json formatted content into the "movie" struct format
func getMovies (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}


// The function below deletes elements from the 
// movies slice by appending all entries after the 
// entery, which is to be deleted, onto the current index  of 
// the  entery to  be deleted and  therefore, well... deleting it.
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r)
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
} 

// The function below returns a single movie from
// the movies slice. The movie to  be returned is  
// picked  by the movie id in the user's  request

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r)
	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
<<<<<<< HEAD
} 

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var myMovie movie
	_=json.NewDecoder(r.Body).Decode(&myMovie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, myMovie)
	json.NewEncoder(w).Encode(myMovie)
}

//The function below deletes a movie and then adds
//and replaces it with a movie, which the user 
//specifies by the movie ID

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item  := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1]...)
			var newMovie movie
			_=json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies  = append(movies, newMovie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}
=======
	json.NewEncoder(w).Encode(item)
} 
>>>>>>> refs/remotes/origin/main

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie(id: "1", isbn: "694201", title: "Pulp non-fiction", director : &director(firstName: "Quentin", lastName: "Tarantino") ))
	movies = append(movies, Movie(id: "2", isbn: "694202", title: "The Scottsman", director : &director(firstName: "Martin", lastName: "Scorsese") ))

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",  deleteMovie).Methods("DELETE")
	
	fmt.Printf("Starting server at Port 8000")
	log.Fatal(http.ListenAndServer(":8000",r))
}	