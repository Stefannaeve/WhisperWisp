package main

import (
	"WhisperWisp/controller"
	"net"
	"net/http"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	var muxServer = http.NewServeMux()
	muxServer.HandleFunc("/wishlist", controller.HandleWishList)
	var listener, error = net.Listen("tcp", ":8080")
	if error != nil {
		os.Exit(1)
	}
	http.Serve(listener, muxServer)
}
