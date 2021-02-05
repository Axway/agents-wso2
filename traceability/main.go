package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Axway/agents-wso2/traceability/pkg/cmd"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"

	// Required Import to setup factory for traceability transport
	_ "github.com/Axway/agent-sdk/pkg/traceability"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

type RequestHeaders struct {
	XHeader1 string `json:"X-Header-1""`
}

type ResponseHeaders struct {
	ContentType   string `json:"Content-Type""`
	ContentLength string `json:"Content-Length"`
}

type Details struct {
	Id              string          `json:"id"`
	Uri             string          `json:"uri"`
	Method          string          `json:"method"`
	SrcHost         string          `json:"srcHost"`
	SrcPort         string          `json:"srcPort"`
	DestHost        string          `json:"destHost"`
	DestPort        string          `json:"destPort"`
	StatusCode      int             `json:"statusCode"`
	RequestHeaders  RequestHeaders  `json:"requestHeaders"`
	ResponseHeaders ResponseHeaders `json:"responseHeaders"`
	RequestBytes    int             `json:"requestBytes"`
	ResponseBytes   int             `json:"responseBytes"`
}

type MessageBody struct {
	ApiName  string  `json:"apiName"`
	Inbound  Details `json:"inbound"`
	Outbound Details `json:"outbound"`
}

func logger(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err.Error())
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			break
		}
		log.Printf(string(message))
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println(string(message))
			break
		}
	}
}

func getTrace(w http.ResponseWriter, r *http.Request) {
	var messageBody MessageBody

	err := json.NewDecoder(r.Body).Decode(&messageBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Printf("%+v", messageBody)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/trace", getTrace).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	// flag.Parse()
	// log.SetFlags(0)

	// upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// http.HandleFunc("/trace", logger)
	// log.Fatal(http.ListenAndServe(*addr, nil))

	//go handleRequests()

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
