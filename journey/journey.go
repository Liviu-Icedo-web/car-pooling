package journey

import (
	"encoding/json"	
	"net/http"
	"log"	
)

type Journey struct {
	Id int `json:"id"`
	Peopel int `json:"people"`	
}

var Journeys = make(map[int] Journey)

func AddJourney(w http.ResponseWriter, r *http.Request) {	
	var journey Journey
	err := json.NewDecoder(r.Body).Decode(&journey)
	if err != nil{
		http.Error(w, "Bad Request", http.StatusBadRequest)
		w.WriteHeader(404)
		return
	}
	if _, ok := Journeys[journey.Id]; ok {
		http.Error(w, "Already exists an user with the same id", http.StatusConflict)
		return
	}else{
		Journeys[journey.Id] = journey
		log.Println("Post Journeys: ", journey)
	}	
}

func GetJourneys(w http.ResponseWriter, r *http.Request) {	
	journeysList := make([]Journey, 0)
	log.Println("All Journeys: ", Journeys)
	log.Println("All Journeys: ",journeysList)
	for _,value := range Journeys {
		journeysList = append(journeysList, value)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(journeysList)
}