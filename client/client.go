package client

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func SendSoap() GetStationBoardResult {

	var resp GetStationBoardResult
	url := fmt.Sprintf("%s%s%s",
		"https://lite.realtime.nationalrail.co.uk",
		"/OpenLDBWS",
		"/ldb11.asmx",
	)
	crs := "SYL"
	rows := 4
	payload := []byte(strings.TrimSpace(`
    <Envelope xmlns="http://www.w3.org/2003/05/soap-envelope">
    <Header>
        <AccessToken xmlns="http://thalesgroup.com/RTTI/2013-11-28/Token/types">
            <TokenValue>` + token + `</TokenValue>
        </AccessToken>
    </Header>
    <Body>
        <GetDepartureBoardRequest xmlns="http://thalesgroup.com/RTTI/2017-10-01/ldb/">
            <numRows>` + strconv.Itoa(rows) + `</numRows>
			<crs>` + crs + `</crs>
			<filterCrs>FEL</filterCrs>
            <filterType>to</filterType>
        </GetDepartureBoardRequest>
    </Body>
</Envelope>`,
	))

	// soapAction := "GetDepartureBoardRequest"

	httpMethod := "POST"

	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return resp
	}

	req.Header.Set("Content-type", "text/xml")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return resp
	}

	result := new(Boardresults)
	err = xml.NewDecoder(res.Body).Decode(result)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		return resp
	}

	resp = result.Body.GetDepartureBoardResponse.GetStationBoardResult
	fmt.Println(result.Body.GetDepartureBoardResponse.GetStationBoardResult)
	return resp

}
