package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "io/ioutil"
)

type Wod struct {
  Id string `json:"id"`
  Title string `json:"title"`
  Desc string `json:"desc"`
  Content string `json:"content"`
}

var Wods = []Wod {
    Wod {Id: "1", Title: "WOD01", Desc: "First day workout of the day", Content: "Run 5km"},
    Wod {Id: "2", Title: "WOD02", Desc: "Second day workout of the day", Content: "Pushing a kettlebell over a long cycle"},
  }

func returnAllWods(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Endpoint Hit: returnAllWods")
  json.NewEncoder(w).Encode(Wods)
}

func returnSingleWod(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  key := vars["id"]
  fmt.Fprintf(w, "Key: " + key)

  for _, wod := range Wods {
    if wod.Id == key {
      json.NewEncoder(w).Encode(wod)
    }
  }
}

func deleteWod(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]
  for index, wod := range Wods {
    if wod.Id == id {
      Wods = append(Wods[:index], Wods[index+1:]...)
    }
  }
}

func createNewWod(w http.ResponseWriter, r *http.Request) {
  reqBody, _ := ioutil.ReadAll(r.Body)
  fmt.Fprintf(w, "%+v", string(reqBody))
  var wod Wod
  json.Unmarshal(reqBody, &wod)
  Wods = append(Wods, wod)
  json.NewEncoder(w).Encode(wod)
}

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome to Home Page")
}

func handleRequests() {
  myRouter := mux.NewRouter().StrictSlash(true)
  myRouter.HandleFunc("/", homePage)
  myRouter.HandleFunc("/all", returnAllWods)
  myRouter.HandleFunc("/wod/{id}", returnSingleWod)
  myRouter.HandleFunc("/wod", createNewWod).Methods("POST")
  myRouter.HandleFunc("/wod/{id}", deleteWod).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8080", myRouter))
}
//func handleRequests() {
//  http.HandleFunc("/", homePage)
//  http.HandleFunc("/wods", returnAllWods)
//  log.Fatal(http.ListenAndServe(":8080", nil))
//}

func main() {
  handleRequests()
}
