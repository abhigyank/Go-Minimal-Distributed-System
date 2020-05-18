package main

import "fmt"
import "flag"


func main() {
	amIMaster := flag.Bool("createMaster", false, "Create this node as master.")
	myPort :=  flag.String("nodePort", "8001", "Port to run this node on.")
	flag.Parse()

	if (*amIMaster) {
		fmt.Println("Creating master node...")
		fmt.Println("Master will listen on port =", *myPort)
	}
}