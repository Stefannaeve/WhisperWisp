package controller

import (
	"WhisperWisp/service"
	"log"
	"net/http"
)

func HandleWishList(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		GetWishList(response)

	}
}

func GetWishList(response http.ResponseWriter) {
	var object = service.GetWishList()
	var numberOfBytes, errorMessage = response.Write([]byte(object))
	if numberOfBytes <= 0 {
		log.Printf("Got no data from service")
	}
	if errorMessage != nil {
		log.Printf("Error message: %s", errorMessage.Error())
	}
}
