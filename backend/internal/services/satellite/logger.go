package satellite

import (
	"log"
)

func logRequestStarted() {
	log.Println("Request Started")
}

func logSendingRequest() {
	log.Println("Sending Request")
}

func logResponseReceived() {
	log.Println("Response Received")
}

func logParsing() {
	log.Println("Parsing")
}

func logFinished() {
	log.Println("Finished")
}
