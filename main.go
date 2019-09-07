package main

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "k8s.io/klog"
)

type Car struct {
    ID        string   `json:"id,omitempty"`
    Model     string   `json:"model,omitempty"`
    Year      string   `json:"lastname,omitempty"`
    Brand    *Brand    `json:"brand,omitempty"`
}

type Brand struct {
    Name  string `json:"city,omitempty"`
}

var cars []Car

func GetCarEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range cars {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Car{})
}

func GetCarsEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(cars)
}

func CreateCarEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var car Car
    _ = json.NewDecoder(req.Body).Decode(&car)
    car.ID = params["id"]
    cars = append(cars, car)
    json.NewEncoder(w).Encode(cars)
}

func DeleteCarEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range cars {
        if item.ID == params["id"] {
            cars = append(cars[:index], cars[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(cars)
}

func main() {
    router := mux.NewRouter()
    cars = append(cars, Car{ID: "1", Model: "SQ5", Brand: &Brand{Name: "Audi"}})
    cars = append(cars, Car{ID: "2", Model: "488 Spider", Brand: &Brand{Name: "Ferrai"}})
    router.HandleFunc("/cars", GetCarsEndpoint).Methods("GET")
    router.HandleFunc("/cars/{id}", GetCarEndpoint).Methods("GET")
    router.HandleFunc("/cars/{id}", CreateCarEndpoint).Methods("POST")
    router.HandleFunc("/cars/{id}", DeleteCarEndpoint).Methods("DELETE")
    klog.Fatal(http.ListenAndServe(":80", router))
}
