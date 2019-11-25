package server

import (
	"encoding/json"
	"fmt"
	"github.com/Mohannad-Zeido/CommuteChecker/client"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

//Init initialises the server and starts listening
func Init() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getBoardByStation).Methods("POST")
	fmt.Println("Server ready and listenng")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getBoardByStation(w http.ResponseWriter, r *http.Request) {
	var newEvent client.RequestParameters

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		_, _ = fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	if err = json.Unmarshal(reqBody, &newEvent); err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println(newEvent)

	if err = json.NewEncoder(w).Encode(
		client.SendSoap(newEvent),
	); err != nil {
		return
	}
}
