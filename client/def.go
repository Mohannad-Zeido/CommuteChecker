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
	Sta                     string          `xml:"sta"`
	Eta                     string          `xml:"eta"`
	Std                     string          `xml:"std"`
	Etd                     string          `xml:"etd"`
	CancelReason            string          `xml:"cancelReason"`
	DelayReason             string          `xml:"delayReason"`
	ServiceID               string          `xml:"serviceID"`
	AdhocAlerts             string          `xml:"adhocAlerts"`
	Operator                string          `xml:"operator"`
	OperatorCode            string          `xml:"operatorCode"`
	Rsid                    string          `xml:"rsid"`
	ServiceType             string          `xml:"serviceType"`
	Length                  string          `xml:"length"`
	Platform                string          `xml:"platform"`
	IsCircularRoute         bool            `xml:"isCircularRoute"`
	IsCancelled             bool            `xml:"isCancelled"`
	FilterLocationCancelled bool            `xml:"filterLocationCancelled"`
	DetachFront             bool            `xml:"detachFront"`
	IsReverseFormation      bool            `xml:"isReverseFormation"`
	Origin                  ServiceLocation `xml:"origin" json:"origin,omitempty"`
	Destination             ServiceLocation `xml:"destination" json:"destination,omitempty"`
	CurrentOrigins          ServiceLocation `xml:"currentOrigins" json:"currentOrigins,omitempty"`
	CurrentDestinations     ServiceLocation `xml:"currentDestinations"`
}

//platform This has special conditions
type ServiceLocation struct {
	LocationName     string `xml:"locationName" json:"location_name,omitempty"`
	Crs              string `xml:"crs" json:"crs,omitempty"`
	Via              string `xml:"via" json:"via,omitempty"`
	FutureChangeTo   string `xml:"futureChangeTo" json:"future_change_to,omitempty"`
	AssocIsCancelled string `xml:"assocIsCancelled" json:"assoc_is_cancelled,omitempty"`
}

type RequestParameters struct {
	FromStation  string `json:"fromStation"`
	ToStation    string `json:"toStation"`
	NumberOfRows string `json:"numberOfRows"`
}
