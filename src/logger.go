package main

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	// file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

}

// func main() {
// 	InfoLogger.Println("Starting the application...")
// 	InfoLogger.Println("Something noteworthy happened")
// 	WarningLogger.Println("There is something you should know about")
// 	ErrorLogger.Println("Something went wrong")
// }

// https://www.honeybadger.io/blog/golang-logging/
// https://appdividend.com/2019/11/28/golang-log-example-log-package-in-go/
