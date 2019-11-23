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
	Dta                     string          `xml:"sta"`
	Eta                     string          `xml:"eta"`
	Dtd                     string          `xml:"std"`
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
	IsCircularRoute         bool            `xml:"isCircularRoute"`
	Platform                bool            `xml:"platform"` //This has special conditions
	IsCancelled             bool            `xml:"isCancelled"`
	FilterLocationCancelled bool            `xml:"filterLocationCancelled"`
	DetachFront             bool            `xml:"detachFront"`
	IsReverseFormation      bool            `xml:"isReverseFormation"`
	Origin                  ServiceLocation `xml:"origin"`
	Destination             ServiceLocation `xml:"destination"`
	CurrentOrigins          ServiceLocation `xml:"currentOrigins"`
	CurrentDestinations     ServiceLocation `xml:"currentDestinations"`
}

type ServiceLocation struct {
	LocationName     string
	Crs              string
	Via              string
	FutureChangeTo   string
	AssocIsCancelled string
}
