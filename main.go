package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Flag variables
var serverPort int
var databaseName string
var databaseHost string
var databasePort int
var databaseProtocol string

// Derived variables
var databaseURL string

func main() {
	flag.IntVar(&serverPort,
		"server-port", 8080,
		"listen for HTTP POST request on this port")
	flag.StringVar(&databaseName,
		"database-name", "mydb",
		"submit data to the InfluxDB database with this name")
	flag.StringVar(&databaseHost,
		"database-host", "localhost",
		"submit data to the InfluxDB on this host")
	flag.IntVar(&databasePort,
		"database-port", 8086,
		"submit data to the InfluxDB on this port")
	flag.StringVar(&databaseProtocol,
		"database-protocol", "http",
		"use this protocol when connecting to InfluxDB API")
	flag.Parse()

	databaseURL = fmt.Sprintf("%s://%s:%d/write?db=%s", databaseProtocol, databaseHost, databasePort, databaseName)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dataBuf := r.Body

		//
		// Maybe decode/transform data here before submitting to InfluxDB
		//

		res, err := http.Post(databaseURL, "text/plain", dataBuf)
		if err != nil {
			log.Print(err)
			return
		}

		// Print the response body for debugging purposes
		responseBody, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Print(err)
			return
		}
		fmt.Printf("%s", responseBody)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil))
}
