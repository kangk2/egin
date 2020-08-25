package lib

import (
	"errors"
	"log"
	"net/http"
)

func Get(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	// new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("new request is fail ")
	}
	// add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	// add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	// http client
	client := &http.Client{}
	log.Printf("Go %s URL : %s \n", http.MethodGet, req.URL.String())
	return client.Do(req)
}
