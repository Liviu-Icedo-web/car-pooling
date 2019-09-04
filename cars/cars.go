package cars
import (
	"encoding/json"	
	"net/http"
	"log"
	util "github.com/car-pooling-challenge/utils"	
)

type Car struct {
	Id int `json:"id"`
	Seats int `json:"seats"`	
}

var Cars = make(map[int] Car)

func AddCar(w http.ResponseWriter, r *http.Request) {	
	var car Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil{
		http.Error(w, "Bad Request", http.StatusBadRequest)
		w.WriteHeader(404)
		return
	}
	if _, ok := Cars[car.Id]; ok {
		http.Error(w, "Already exists an user with the same id", http.StatusConflict)
		return
	}else{
		Cars[car.Id] = car
		log.Println("Post Cars: ", car)
	}	
}

func GetCars(w http.ResponseWriter, r *http.Request) {	
	carsList := make([]Car , 0)
	log.Println("All Cars: ", Cars)
	log.Println("All Cars: ",carsList)
	for _,value := range Cars {
		carsList = append(carsList, value)
	}
	util.Respond(w,carsList)
    
	
}