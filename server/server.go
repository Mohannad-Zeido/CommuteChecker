package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Mohannad-Zeido/CommuteChecker/client"
	"github.com/gorilla/mux"
)

type requestParameters struct {
	FromStation  string `json:"fromStation"`
	ToStation    string `json:"toStation"`
	NumberOfRows string `json:"numberOfRows"`
}

func Init() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getBoardByStation).Methods("POST")
	fmt.Println("Server ready and listenng")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getBoardByStation(w http.ResponseWriter, r *http.Request) {
	var newEvent requestParameters

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		_, _ = fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	if err = json.Unmarshal(reqBody, &newEvent); err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println(newEvent)

	//todo: Place this somewhere else
	newEvent.FromStation = strings.ToUpper(newEvent.FromStation)
	newEvent.ToStation = strings.ToUpper(newEvent.ToStation)

	if err = json.NewEncoder(w).Encode(
		client.SendSoap(newEvent.FromStation, newEvent.ToStation, newEvent.NumberOfRows),
	); err != nil {
		return
	}
}
