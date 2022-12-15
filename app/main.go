package main

import (
	"log"
	"net/http"
	"os"
)

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return hostname
}

const DefaultPort = "8080"

func getServerPort() string {
	port := os.Getenv("SERVER_PORT")
	if port != "" {
		return port
	}
	return DefaultPort
}

func pingHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("ping received from " + request.RemoteAddr)
	writer.Write([]byte("pong"))
}

func hostnameHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("hostname received from " + request.RemoteAddr)
	writer.Write([]byte(getHostname()))
}

func echoHandler(writer http.ResponseWriter, request *http.Request) {

	log.Println("Echoing back request made to " + request.URL.Path + " to client (" + request.RemoteAddr + ")")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")
	request.Write(writer)
}

func main() {

	log.Println("starting server, listening on port " + getServerPort())

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/hostname", hostnameHandler)
	http.ListenAndServe(":"+getServerPort(), nil)
}
