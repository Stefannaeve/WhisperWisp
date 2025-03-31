package main

import (
	"WhisperWisp/controller"
	"log"
	"net"
	"net/http"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	log.Println("Application startin...")
	var muxServer = http.NewServeMux()
	muxServer.HandleFunc("/wishlist", controller.HandleWishList)
	var listener, errorMessage = net.Listen("tcp", ":8080")
	if errorMessage != nil {
		os.Exit(1)
	}
	log.Println("Starting server...")
	errorMessage = http.Serve(listener, muxServer)
	if errorMessage != nil {
		log.Println("ErrorMessage: ", errorMessage.Error())
	}
}
