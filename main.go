package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	. "github.com/car-pooling-challenge/users"
	Cars  "github.com/car-pooling-challenge/cars"
	Journeys "github.com/car-pooling-challenge/journey"
	"strconv"
)

var Users map[int] User


func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		//http.Error(w, "Bad Request", http.StatusBadRequest)
		w.WriteHeader(404)
		return
	}
	if _, ok := Users[user.Id]; ok {
		http.Error(w, "Already exists an user with the same id", http.StatusConflict)
		return
	}else{
		Users[user.Id] = user
	}
	log.Println("User: ", user, "added")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userList := make([]User, 0)
	log.Println("users: ",Users)
	log.Println("userList: ",userList)
	for _,value := range Users {
		userList = append(userList, value)
	}
	
	json.NewEncoder(w).Encode(userList)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if user, ok := Users[id]; ok {
		json.NewEncoder(w).Encode(user)
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if user, ok := Users[id]; ok {
		delete(Users, id)
		log.Println("User: ", user, "removed")
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func CheckStatus(w http.ResponseWriter, r *http.Request){
	 resp,err := http.Get("http://localhost:8005/users")
	 if( err == nil) && (resp.StatusCode == 200){
		 w.WriteHeader(200)
		 return
	 }else{
		 w.WriteHeader(500)
	 }
	 
}

func main(){
	Users = make(map[int]User)
	log.Println("Default users: ",Users)
	router := mux.NewRouter()
	router.HandleFunc("/cars", Cars.AddCar).Methods("POST")
	router.HandleFunc("/cars", Cars.GetCars).Methods("GET")
	router.HandleFunc("/journey", Journeys.AddJourney).Methods("POST")
	router.HandleFunc("/journey", Journeys.GetJourneys).Methods("GET")
	router.HandleFunc("/status", CheckStatus).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")	
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	log.Println("Listening on http://localhost:8005")
	log.Fatal(http.ListenAndServe(":8005",router))

}