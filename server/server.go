package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../client"
	"github.com/gorilla/mux"
)

type requestParameters struct {
	FromStation  string `json:"fromStation"`
	ToStation    string `json:"toStation"`
	NumberOfRows int    `json:"numberOfRows"`
}

type getStationBoardResult struct {
	GeneratedAt          string        `json:"generatedAt"`
	LocationName         string        `json:"locationName"`
	Crs                  string        `json:"crs"`
	FilterLocationName   string        `json:"filterLocationName"`
	Filtercrs            string        `json:"filtercrs"`
	FilterType           string        `json:"filterType"`
	NrccMessages         string        `json:"nrccMessages"`
	PlatformAvailable    bool          `json:"platformAvailable"`
	AreServicesAvailable bool          `json:"areServicesAvailable"`
	TrainServices        trainServices `json:"trainServices"`
}

type trainServices struct {
	Service []service `json:"service"`
}

type service struct {
	Std string `json:"std"`
}

func Init() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getBoardByStation).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getBoardByStation(w http.ResponseWriter, r *http.Request) {
	var newEvent requestParameters
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	w.WriteHeader(http.StatusOK)
	fmt.Println(newEvent)
	json.NewEncoder(w).Encode(client.SendSoap())
}
