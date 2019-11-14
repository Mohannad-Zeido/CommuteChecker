package client

type boardResult struct {
	Body struct {
		GetDepartureBoardResponse getDepartureBoardResponse `xml:"GetDepartureBoardResponse"`
	}
}

type getDepartureBoardResponse struct {
	StationBoardResult StationBoardResult `xml:"GetStationBoardResult"`
}

type StationBoardResult struct {
	RequestTime          string        `xml:"generatedAt"`
	FromStationName      string        `xml:"locationName"`
	FromStationCode      string        `xml:"crs"`
	ToStationName        string        `xml:"filterLocationName"`
	ToStationCode        string        `xml:"filtercrs"`
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
	ExpectedTime string `xml:"std"`
}
