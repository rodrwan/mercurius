package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rodrwan/mercurius/server"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	port = kingpin.Flag("port", "Server port.").Short('p').String()
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	maxProcs := runtime.NumCPU()
	fmt.Printf("Available CPU %d\n", maxProcs)
	runtime.GOMAXPROCS(maxProcs)

	maxQueueSize, _ := strconv.Atoi(os.Getenv("MAX_QUEUE_SIZE"))
	maxWorkers, _ := strconv.Atoi(os.Getenv("MAX_WORKERS"))

	kingpin.Parse()
	var serverPort string
	fmt.Printf("Setting JobQueue to handle %d jobs\n", maxQueueSize)
	server.JobQueue = make(chan server.Job, maxQueueSize)

	// Start the dispatcher.
	dispatcher := server.NewDispatcher(server.JobQueue, maxWorkers)
	dispatcher.Run()

	serverPort = fmt.Sprintf(":%s", *port)
	if p := os.Getenv("PORT"); p != "" {
		serverPort = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}

	server.API(serverPort)
}
