package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func Index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I am the index!\n")
}

func main() {

	var port int
	var err error

	port_env, ok := os.LookupEnv("PORT")

	if !ok {
		fmt.Printf("Using default port %d", 8000)
	} else {
		port, err = strconv.Atoi(port_env)
		if err != nil {
			log.Fatal(err)
		}
	}

	router := mux.NewRouter()
	router.HandleFunc("/hello", HelloServer)
	router.HandleFunc("/", Index)

	srv := http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Starting web service on port %d", port)
	log.Fatal(srv.ListenAndServe())
}
