package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

// Method ...
func Method(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request incoming...")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// Scraping response
	vars := mux.Vars(r)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, maxLength))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	taskType := vars["type"]

	if taskType != "transactions" {
		resp := APICall(body, taskType)
		fmt.Println(resp)
		w.WriteHeader(http.StatusOK)
		return
	}

	name := RandStringRunes(5)
	fmt.Printf("Job %s created\n", name)
	// Create Job and push the work onto the jobQueue.
	job := Job{Name: name, Payload: body, EndPoint: taskType}
	JobQueue <- job
	fmt.Println("Enqueue job")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
