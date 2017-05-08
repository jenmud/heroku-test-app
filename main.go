package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
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

	http.HandleFunc("/hello", HelloServer)

	fmt.Printf("Starting web service on port %d", port)
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", port),
			nil,
		),
	)
}
