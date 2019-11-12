package client

type Boardresults struct {
	Body struct {
		GetDepartureBoardResponse GetDepartureBoardResponse `xml:"GetDepartureBoardResponse"`
	}
}

type GetDepartureBoardResponse struct {
	GetStationBoardResult GetStationBoardResult `xml:"GetStationBoardResult"`
}
type GetStationBoardResult struct {
	GeneratedAt          string        `xml:"generatedAt"`
	LocationName         string        `xml:"locationName"`
	Crs                  string        `xml:"crs"`
	FilterLocationName   string        `xml:"filterLocationName"`
	Filtercrs            string        `xml:"filtercrs"`
	FilterType           string        `xml:"filterType"`
	NrccMessages         string        `xml:"nrccMessages"`
	PlatformAvailable    bool          `xml:"platformAvailable"`
	AreServicesAvailable bool          `xml:"areServicesAvailable"`
	TrainServices        TrainServices `xml:"trainServices"`
}

type TrainServices struct {
	Service []Service `xml:"service"`
}

type Service struct {
	Std string `xml:"std"`
}
