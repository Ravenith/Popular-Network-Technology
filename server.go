package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	//GET
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(writer, "This is server MUX")
	})

	//POST
	mux.HandleFunc("/post", func(writer http.ResponseWriter, request *http.Request) {
		//get data from client's request
		data, err := io.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}
		//unbox data from client
		var requestData map[string]string
		err = json.Unmarshal(data, &requestData)
		if err != nil {
			panic(err)
		}
		//print result
		fmt.Println("Received data: ", requestData)
		fmt.Println(writer, "Data successfully received by server")

	})

	server := http.Server{
		Addr:    "localhost:4444",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
