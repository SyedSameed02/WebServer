package main

import ( 
	
	"fmt"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"	
	"strconv"
 )

// "math/random"
// 	

// func helloHandler(w http.ResponseWriter , r *http.Request){

// 	if r.URL.Path != "/hello"{
// 	http.Error(w,"404 not found",http.StatusNotFound)
// 	return
// 	}
// 	if r.Method != "GET" {
// 		 http.Error(w,"method not supported",http.StatusNotFound)
// 		 return
// 	}
// 	fmt.Fprintf(w,"hello!")
// }

// func formHandler(w http.ResponseWriter, r *http.Request){

// 	if err := r.ParseForm(); err != nil {
// 		fmt.Fprintf(w,"parseForm() err: %v",err)
// 		return
// 	}
// 	fmt.Fprintf(w,"Post request successful")
// 	name := r.FormValue("name")
// 	address := r.FormValue("address")

// 	fmt.Fprintf(w,"name = %s\n",name)
// 	fmt.Fprintf(w,"address = %s\n",address)
// }



type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}


type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie 


func getMovies(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	for _,item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func delteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	for idx,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:idx],movies... )
			break
		}
	}
}

func addMovie( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(len(movies)+1)

	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	for idx, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:idx],movies... )
			
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]

			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
		
	}
}

func main () {

	r := mux.NewRouter()

	//static data

	movies = append(movies, Movie{ID: "1",Isbn: "34623",Title: "it",Director: &Director{Firstname: "raja",Lastname: "mouli"}})

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}",getMovie).Methods("GET")
	r.HandleFunc("/delete/{id}",delteMovie).Methods("DELETE")
	r.HandleFunc("/movies",addMovie).Methods("POST")
	r.HandleFunc("/movies",updateMovie).Methods("PUT")


	fmt.Print("Server is started at port 8080\n")

	log.Fatal(http.ListenAndServe(":8080",r))

}


