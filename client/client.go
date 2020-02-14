package client

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"github.com/Mohannad-Zeido/CommuteChecker/authentication"
	"log"
	"net/http"
	"strings"
)

//SendSoap sends soap message to the OpenLDBWS api
func SendSoap(params RequestParameters) StationBoardResult {
	var resp StationBoardResult
	token, _ := authentication.GetToken()

	params.FromStation = strings.ToUpper(params.FromStation)
	params.ToStation = strings.ToUpper(params.ToStation)

	url := fmt.Sprintf("%s%s%s",
		"https://lite.realtime.nationalrail.co.uk", "/OpenLDBWS", "/ldb11.asmx",
	)

	payload := []byte(strings.TrimSpace(`
	   <Envelope xmlns="http://www.w3.org/2003/05/soap-envelope">
	   <Header>
	       <AccessToken xmlns="http://thalesgroup.com/RTTI/2013-11-28/Token/types">
	           <TokenValue>` + token + `</TokenValue>
	       </AccessToken>
	   </Header>
	   <Body>
	       <GetDepartureBoardRequest xmlns="http://thalesgroup.com/RTTI/2017-10-01/ldb/">
	           <numRows>` + params.NumberOfRows + `</numRows>
				<crs>` + params.FromStation + `</crs>
				<filterCrs>` + params.ToStation + `</filterCrs>
	       </GetDepartureBoardRequest>
	   </Body>
	</Envelope>`,
	))
	httpMethod := "POST"
	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(payload))

	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return resp
	}

	req.Header.Set("Content-type", "text/xml")

	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{},
		},
	}

	res, err := c.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return resp
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Fatal("Error on close. ", err.Error())
			return
		}
	}()

	result := new(boardResult)

	if err = xml.NewDecoder(res.Body).Decode(result); err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		return resp
	}

	resp = result.Body.GetDepartureBoardResponse.StationBoardResult

	fmt.Println("resp")
	fmt.Println(resp)

	return resp
}
