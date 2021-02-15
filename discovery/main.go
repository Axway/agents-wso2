package main

import (
	"fmt"
	"time"

	"github.com/Axway/agents-wso2/discovery/pkg/cmd"
)

func main() {

	
	go forever()
	select {} // block forever
}

func forever() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		// os.Exit(1)
	}
	for {
		// fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second)
	}
}
