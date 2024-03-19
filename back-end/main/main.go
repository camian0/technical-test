package main

import (
	"fmt"
	"myapp.com/enron/mbox"
	server2 "myapp.com/enron/server"
)

func main() {
	mbox.ConvertirMbox()
	mbox.ReadMbox()

	mux := Routes()
	server := server2.CreateServer(mux)
	fmt.Println("Servidor corriendo...")
	fmt.Println("Puerto: 3000")
	server.Run()
}
