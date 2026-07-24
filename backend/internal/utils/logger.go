package utils

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "[MWANGAZA] ", log.LstdFlags|log.Lshortfile)

func Info(msg string) {
	logger.Println("INFO: " + msg)
}

func Error(msg string, err error) {
	logger.Println("ERROR: "+msg, err)
}

func Fatal(msg string, err error) {
	logger.Fatal("FATAL: "+msg, err)
}