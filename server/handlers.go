package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var maxLength int64 = 1048576

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStringRunes ...
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// Index ..
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Categorizer!")
}

// Method ...
func Method(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request incoming...")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// Scraping response
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, maxLength))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	name := RandStringRunes(5)
	fmt.Printf("Job %s created\n", name)
	// Create Job and push the work onto the jobQueue.
	job := Job{Name: name, Payload: body}
	JobQueue <- job
	fmt.Println("Enqueue job")
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	// w.Write(stream)
}

// PayloadHandler ...
func PayloadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("running payload")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, maxLength))
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name := RandStringRunes(5)
	// let's create a job with the payload
	job := Job{Name: name, Payload: body}
	JobQueue <- job
	w.WriteHeader(http.StatusOK)
}
