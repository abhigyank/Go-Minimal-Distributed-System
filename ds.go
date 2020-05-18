package main

import (
	"fmt"
	"os"
)
import "flag"
import "net"


var ipAddress string = "127.0.0.1"

func createMaster (port string) {
	/* Start listening on port */
	ln, err := net.Listen("tcp", fmt.Sprint(":" + port))
	if err != nil {
		fmt.Println("Error occured in creating master: ", err)
		os.Exit(0)
	}
	fmt.Println("Master is now listening on port =", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error occured in accepting connection: ", err)
		} else {
			fmt.Println("Received request to master.")
			// Code here for add node to cluster
			conn.Close()
		}
	}
}

func main() {
	amIMaster := flag.Bool("createMaster", false, "Create this node as master.")
	myPort :=  flag.String("nodePort", "8001", "Port to run this node on.")
	flag.Parse()

	if (*amIMaster) {
		fmt.Println("Creating master node...")
		fmt.Println("Master will listen on port =", *myPort)
		createMaster(*myPort)
	}

}