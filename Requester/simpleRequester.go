package Requester

import (
	"bytes"
	"github.com/gowebspider/utensils/Head"
	"github.com/saintfish/chardet"
	"github.com/wonderivan/logger"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// SimplePcFetch SimpleFetch is Network requester, It relies on net / HTTP packets to send network requests
// * fake User Agent
// * AutoEncode
//func SimplePcFetch(method, url string, body io.Reader) ([]byte, error) {
//	client := &http.Client{}
//	req, err := http.NewRequest(method, url, body)
//	if err != nil {
//		return nil, fmt.Errorf(`NewRequest Error:%#v`, err)
//	}
//	Head.SetRandomPcUserAgent(req)
//	res, err := client.Do(req)
//	if err != nil {
//		return nil, fmt.Errorf(` Client Error:%#v`, err)
//	}
//	defer res.Body.Close()
//	if res.StatusCode != http.StatusOK {
//		return nil, fmt.Errorf(` Get invalid status code %s while scraping %s`, res.Status, url)
//	}
//	// Detection encode，eg: UTF8 GBK...
//	bodyReader := bufio.NewReader(res.Body)
//	PageEncode := Encoder.DetectionEncode(bodyReader)
//	// recoding
//	encodedReader := transform.NewReader(bodyReader, PageEncode.NewDecoder())
//	return ioutil.ReadAll(encodedReader)
//}
func fixCharset(r *http.Response, detectCharset bool, defaultEncoding string) ([]byte, error) {
	Body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error(`ReadAll Error %v`, err)
	}
	if defaultEncoding != "" {
		tmpBody, err := encodeBytes(Body, "text/plain; charset="+defaultEncoding)
		if err != nil {
			return nil, nil
		}
		Body = tmpBody
		logger.Error(`defaultEncoding  Error %v`, err)
	}
	contentType := strings.ToLower(r.Header.Get("Content-Type"))

	if strings.Contains(contentType, "image/") ||
		strings.Contains(contentType, "video/") ||
		strings.Contains(contentType, "audio/") ||
		strings.Contains(contentType, "font/") {
		// These MIME types should not have textual data.

		logger.Error(`strings.Contains  Error %v`, err)
	}

	if !strings.Contains(contentType, "charset") {
		if !detectCharset {
			logger.Error(`!strings.Contains  Error %v`, err)
		}
		d := chardet.NewTextDetector()
		r, err := d.DetectBest(Body)
		if err != nil {
			logger.Error(`DetectBest  Error %v`, err)
		}
		contentType = "text/plain; charset=" + r.Charset
	}
	if strings.Contains(contentType, "utf-8") || strings.Contains(contentType, "utf8") {
		logger.Warn(`strings.Contains  not is utf8 %v`, err)
	}
	return encodeBytes(Body, contentType)
}

func encodeBytes(b []byte, contentType string) ([]byte, error) {
	r, err := charset.NewReader(bytes.NewReader(b), contentType)
	if err != nil {
		logger.Error(`charset.NewReader %v`, err)
	}
	return ioutil.ReadAll(r)
}

func SimplePcFetch(method, url string, body io.Reader) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		logger.Error(`http.NewRequest Err %v`, err)
	}
	Head.SetRandomPcUserAgent(req)
	res, err := client.Do(req)
	if err != nil {
		logger.Error(`client.Do Err %v`, err)
	}
	defer res.Body.Close()
	return fixCharset(res, true, ``)
}
