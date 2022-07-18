package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// handler handles an incoming HTTP request
func handler(writer http.ResponseWriter, request *http.Request) {
	keySet, err := makeSet()

	if err != nil {
		log.Fatal(err)
	}

	jsonWebKeySet, _ := json.MarshalIndent(keySet, "", "  ")
	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, "%v", string(jsonWebKeySet))
}
