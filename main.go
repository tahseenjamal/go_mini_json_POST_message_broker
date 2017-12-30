package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func jsonrequest(url string) {

	//infinte loop to keep accepting the data from queue

	for {

		// blocking queue. This will also pause the function at this place until the data is received
		// so in a way good as infinite loop wouldn't consume CPU

		newData := <-jsonPipe

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(newData))

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {

			panic(err)

		}

		defer resp.Body.Close()

	}

}

func handler(resp http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()

	//read body for json data
	msg, _ := ioutil.ReadAll(req.Body)

	//pass the data to the message queue for the threads to read and call the target URL
	jsonPipe <- msg

	fmt.Fprintf(resp, "200")

}

func showQueue(resp http.ResponseWriter, req *http.Request) {

	responseString := fmt.Sprintf("Total queue buffer = %v", len(jsonPipe))

	fmt.Fprintf(resp, responseString)

}

func returnParameters() (int, string, string) {

	var threads int

	var port, targetURL string

	threads, _ = strconv.Atoi(os.Args[1])

	port = fmt.Sprintf(":%v", os.Args[2])

	targetURL = os.Args[3]

	return threads, port, targetURL

}

var jsonPipe chan []byte

var url string

func main() {

	var threads int

	var port string

	threads, port, url = returnParameters()

	jsonPipe = make(chan []byte)

	for loop := 0; loop < threads; loop++ {

		go jsonrequest(url)

	}

	http.HandleFunc("/", handler)

	http.HandleFunc("/queue", showQueue)

	log.Fatal(http.ListenAndServe(port, nil))

}
