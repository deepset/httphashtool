package script

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const REQUEST_TIMEOUT = 5

// getHashFromURL make http get request to url and get md5 hash from the response body
func GetHashFromURL(requestUrl string) (string, error) {

	var response string
	// url validation
	url, err := url.Parse(requestUrl)
	if err != nil {
		return response, err
	}

	//append scheme if empty
	if url.Scheme == "" {
		requestUrl = "http://" + requestUrl
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: REQUEST_TIMEOUT * time.Second,
	}

	// http get request
	getResponse, err := client.Get(requestUrl)
	if err != nil {
		return response, err
	}
	defer getResponse.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(getResponse.Body)
	if err != nil {
		return response, err
	}

	// Convert the body to type string
	text := string(body)
	// Generating the md5 hash
	hash := md5.Sum([]byte(text))

	// creating final response
	response = fmt.Sprintf("%s %x", requestUrl, hash)

	return response, nil
}
