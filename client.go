package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	sendData()
}

func receivedData() {
	//dial http address that server listening
	resp, err := http.Get("http://localhost:4444")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//read data from server
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//unbox & print result
	fmt.Println("Message from server: ", string(data))
}

func sendData() {
	//variable data
	data := map[string]string{
		"Message": "Hello World",
	}
	//box
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	//send data that boxed
	//using buffer
	resp, err := http.Post("http://localhost:4444/post", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	//close
	defer resp.Body.Close()

	//check server success received data/not
	//read data
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//print result
	fmt.Println("Server response: ", string(body))

}
