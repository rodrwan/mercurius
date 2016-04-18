package server

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Job holds the attributes needed to perform unit of work.
type Job struct {
	Name    string
	Payload []byte
}

// CategoryCall ...
func (j *Job) CategoryCall() []byte {
	fmt.Println("Running task")
	body := j.Payload
	categoryURL := os.Getenv("CATEGORY_SERVICE_URL")
	categoryBearer := fmt.Sprintf("Token token=%s", os.Getenv("CATEGORY_BEARER"))
	fmt.Printf("URL: %s\n", categoryURL)
	fmt.Printf("BEARER: %s\n", categoryBearer)
	resp := httpClient(categoryURL, categoryBearer, body)
	return resp
}

// APICall ...
func APICall(body []byte) string {
	fmt.Println("Running task")
	apiURL := os.Getenv("API_URL")
	apiBearer := fmt.Sprintf("Token token=%s", os.Getenv("API_BEARER"))
	resp := httpClient(apiURL, apiBearer, body)
	return string(resp)
}

func httpClient(URL string, bearer string, body []byte) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", URL, bytes.NewReader(body))
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)
	defer resp.Body.Close()
	return streamToByte(resp.Body)
}

func streamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

// JobQueue ...
var JobQueue chan Job
