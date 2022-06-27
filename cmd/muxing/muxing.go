package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", getName).Methods("GET")
	router.HandleFunc("/bad", processBadRequest).Methods("GET")
	router.HandleFunc("/data", postData).Methods("POST")
	router.HandleFunc("/headers", getHeaders).Methods("POST")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func getName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s!", vars["PARAM"])
}

func processBadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func postData(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write([]byte("I got message:\n" + string(body)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func getHeaders(w http.ResponseWriter, r *http.Request) {
	a := r.Header.Get("A")
	numberA, _ := strconv.Atoi(a)

	b := r.Header.Get("B")
	numberB, _ := strconv.Atoi(b)

	w.Header().Set("a+b", strconv.Itoa(numberA+numberB))
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
