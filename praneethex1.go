package main

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
)

func init () {
	movies = []Movie{
		Movie{
		Id: 1,
		Title: "The Terminal",
		Year: 2004,
		Genre: "Drama",
		Cast: []string{"Tom Hanks","Caterine Zeta Jones"},
		Ratings: []float64{7.3, 7.4},
		},
		{
			Id: 2,
			Title: "12 Angry Men",
			Year: 1957,
			Genre: "Drama",
			Cast: []string{"Henry Fonda,","Lee J Cobb,","Martin Balsam"},
			Ratings: []float64{8.9, 9.0, 8.8, 9.3},
		},
	}

}

type Delrating struct{
	Id int `json:"id"`
}

type Rating struct{
	Id int `json:"id"`
	Values []float64 `json:"values"`
}

type Movie struct{
	Id int `json:"Id"`
	Title string `json:"Title"`
	Year int `json:"Year"`
	Cast []string `json:"Cast"`
	Genre string `json:"Genre"`
	Ratings []float64 `json:"Ratings"`
}

var movies []Movie


func handleRequests() {
	C := mux.NewRouter().StrictSlash(true)
	C.HandleFunc("/", page)
	C.HandleFunc("/createPage", createNew)
	C.HandleFunc("/updatePage", updateRatings)
	C.HandleFunc("/deletePage", deleteMovie)
	log.Fatal(http.ListenAndServe(":10000", C))
	return
}

func page(w http.ResponseWriter, r *http.Request){

	reqBody, _ := ioutil.ReadAll(r.Body)
   	new := make([]string, 1)
  	json.Unmarshal(reqBody, &new)
	json.NewEncoder(w).Encode(movies)
	return
}

func createNew(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    var mov Movie 
    json.Unmarshal(reqBody, &mov)
    movies = append(movies, mov)
	json.NewEncoder(w).Encode(movies)
	return
}

func updateRatings(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    var rating Rating 
    json.Unmarshal(reqBody, &rating)
    for j:=0; j<len(movies); j++ {
		if movies[j].Id == rating.Id {
			movies[j].Ratings = append(movies[j].Ratings, rating.Values...)
			json.NewEncoder(w).Encode(movies[j])
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
    var deleteid Delrating 
    json.Unmarshal(reqBody, &deleteid)
	for n := 0; n < len(movies)-1; {
		if movies[n].Id == deleteid.Id { 
	movies = append(movies[:n], movies[n+1:]...)
		}
	}
	json.NewEncoder(w).Encode(movies)
return
}


func main () {
			
handleRequests()
	
}