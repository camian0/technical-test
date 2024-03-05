package main

import (
	"fmt"
	server2 "myapp.com/enron/server"
)

func main() {
	mux := Routes()
	server := server2.CreateServer(mux)
	fmt.Println("Servidor corriendo...")
	fmt.Println("Puerto: 3000")
	server.Run()
}
