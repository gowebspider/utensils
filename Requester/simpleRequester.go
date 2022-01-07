package Requester

import (
	"bufio"
	"bytes"
	"github.com/gowebspider/utensils/Head"
	"github.com/saintfish/chardet"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

type ReqArgs struct {
	method, url  string
	header, data map[string]string
}

func fixCharset(r *http.Response, detectCharset bool, defaultEncoding string) ([]byte, error) {
	Body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf(`ReadAll Error %v`, err)
	}
	if defaultEncoding != "" {
		tmpBody, err := encodeBytes(Body, "text/plain; charset="+defaultEncoding)
		if err != nil {
			return nil, nil
		}
		Body = tmpBody
		logrus.Errorf(`defaultEncoding  Error %v`, err)
	}
	contentType := strings.ToLower(r.Header.Get("Content-Type"))

	if strings.Contains(contentType, "image/") ||
		strings.Contains(contentType, "video/") ||
		strings.Contains(contentType, "audio/") ||
		strings.Contains(contentType, "font/") {
		// These MIME types should not have textual data.

		logrus.Errorf(`strings.Contains  Error %v`, err)
	}

	if !strings.Contains(contentType, "charset") {
		if !detectCharset {
			logrus.Errorf(`!strings.Contains  Error %v`, err)
		}
		d := chardet.NewTextDetector()
		r, err := d.DetectBest(Body)
		if err != nil {
			logrus.Errorf(`DetectBest  Error %v`, err)
		}
		contentType = "text/plain; charset=" + r.Charset
	}
	if strings.Contains(contentType, "utf-8") || strings.Contains(contentType, "utf8") {
		logrus.Warnf(`strings.Contains  not is utf8 %v`, err)
	}
	return encodeBytes(Body, contentType)
}

func encodeBytes(b []byte, contentType string) ([]byte, error) {
	r, err := charset.NewReader(bytes.NewReader(b), contentType)
	if err != nil {
		logrus.Errorf(`charset.NewReader %v`, err)
	}
	return ioutil.ReadAll(r)
}

// DetectionEncode is a Code detector
// It will read 1024 bytes to implement coding detection
func DetectionEncode(r *bufio.Reader) encoding.Encoding {
	peek, err := r.Peek(1024)
	if err != nil {
		log.Printf(`DetectionEncode Error:%#v`, err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(peek, "")
	return e
}

//NewReqArgs Initialize
func NewReqArgs(method, url string, header, data map[string]string) *ReqArgs {
	return &ReqArgs{
		method: method,
		url:    url,
		header: header,
		data:   data,
	}
}

// SimpleScrape based on net/http Client Encapsulation to requests
func (r *ReqArgs) SimpleScrape() ([]byte, error) {
	//Initialize client
	client := &http.Client{}

	// set payload
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for fieldname, value := range r.data {
		_ = writer.WriteField(fieldname, value)
	}
	err := writer.Close()

	// Construct request line
	req, err := http.NewRequest(r.method, r.url, payload)
	if err != nil {
		logrus.Warnln(err)
		return nil, nil
	}

	// set header
	for k, v := range r.header {
		req.Header.Add(k, v)
	}
	// set Random UserAgent
	req.Header.Set(`userAgent`, Head.RandomAllUserAgent())

	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		logrus.Warnln(err)
		return nil, nil
	}
	if res.StatusCode != 200 {
		logrus.Warnf(`response StatusCode not is 200 !!!, response StatusCode:%d`, res.StatusCode)
	}

	defer res.Body.Close()

	//return fixCharset(res, true, ``)
	return fixCharset(res, true, ``)
}
