package main

import (
	"flag"
	"log"
	"net/http"

	// Required Import to setup factory for traceability transport
	_ "github.com/Axway/agent-sdk/pkg/traceability"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

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

func main() {
	// flag.Parse()
	// log.SetFlags(0)

	// upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// http.HandleFunc("/trace", logger)
	// log.Fatal(http.ListenAndServe(*addr, nil))

	// if err := cmd.RootCmd.Execute(); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
}
