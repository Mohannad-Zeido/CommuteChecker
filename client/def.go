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
	RequestTime          string        `xml:"generatedAt" json:"request_time,omitempty"`
	FromStationName      string        `xml:"locationName" json:"from_station_name,omitempty"`
	FromStationCode      string        `xml:"crs" json:"from_station_code,omitempty"`
	ToStationName        string        `xml:"filterLocationName" json:"to_station_name,omitempty"`
	ToStationCode        string        `xml:"filtercrs" json:"to_station_code,omitempty"`
	FilterType           string        `xml:"filterType" json:"filter_type,omitempty"`
	NrccMessages         string        `xml:"nrccMessages" json:"nrcc_messages,omitempty"`
	PlatformAvailable    bool          `xml:"platformAvailable" json:"platform_available,omitempty"`
	AreServicesAvailable bool          `xml:"areServicesAvailable" json:"are_services_available,omitempty"`
	TrainServices        TrainServices `xml:"trainServices" json:"train_services,omitempty"`
}

type TrainServices struct {
	Service []Service `xml:"service"`
}

type Service struct {
	ArrivalTime             string          `xml:"sta" json:"arrival_time,omitempty"`
	EstimatedArrivalTime    string          `xml:"eta" json:"estimated_arrival_time,omitempty"`
	DepartureTime           string          `xml:"std" json:"departure_time,omitempty"`
	EstimatedDepartureTime  string          `xml:"etd" json:"estimated_departure_time,omitempty"`
	CancelReason            string          `xml:"cancelReason" json:"cancel_reason,omitempty"`
	DelayReason             string          `xml:"delayReason" json:"delay_reason,omitempty"`
	ServiceID               string          `xml:"serviceID" json:"service_id,omitempty"`
	AdhocAlerts             string          `xml:"adhocAlerts" json:"adhoc_alerts,omitempty"`
	Operator                string          `xml:"operator" json:"operator,omitempty"`
	OperatorCode            string          `xml:"operatorCode" json:"operator_code,omitempty"`
	Rsid                    string          `xml:"rsid" json:"rsid,omitempty"`
	ServiceType             string          `xml:"serviceType" json:"service_type,omitempty"`
	Length                  string          `xml:"length" json:"length,omitempty"`
	Platform                string          `xml:"platform" json:"platform,omitempty"`
	IsCircularRoute         bool            `xml:"isCircularRoute" json:"is_circular_route,omitempty"`
	IsCancelled             bool            `xml:"isCancelled" json:"is_cancelled,omitempty"`
	FilterLocationCancelled bool            `xml:"filterLocationCancelled" json:"filter_location_cancelled,omitempty"`
	DetachFront             bool            `xml:"detachFront" json:"detach_front,omitempty"`
	IsReverseFormation      bool            `xml:"isReverseFormation" json:"is_reverse_formation,omitempty"`
	Origin                  ServiceLocation `xml:"origin" json:"origin,omitempty"`
	Destination             ServiceLocation `xml:"destination" json:"destination,omitempty"`
	CurrentOrigins          ServiceLocation `xml:"currentOrigins" json:"current_origins,omitempty"`
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
