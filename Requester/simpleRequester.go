package Requester

import (
	"bufio"
	"fmt"
	"github.com/gowebspider/utensils/Encoder"
	"github.com/gowebspider/utensils/Head"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

// SimpleFetch is Network requester, It relies on net / HTTP packets to send network requests
// * fake User Agent
// * AutoEncode
func SimplePcFetch(method, url string, params map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf(`NewRequest Error:%#v`, err)
	}
	Head.SetRandomPcUserAgent(req)
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf(` Client Error:%#v`, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(` Get invalid status code %s while scraping %s`, res.Status, url)
	}
	// Detection encode，eg: UTF8 GBK...
	bodyReader := bufio.NewReader(res.Body)
	PageEncode := Encoder.DetectionEncode(bodyReader)
	// recoding
	encodedReader := transform.NewReader(bodyReader, PageEncode.NewDecoder())
	return ioutil.ReadAll(encodedReader)
}
