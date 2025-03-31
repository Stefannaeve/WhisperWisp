package controller

import (
	"log"
	"net/http"
)

func HandleWishList(response http.ResponseWriter, request *http.Request) {
	var object = "{wish: big giant dick}"
	if request.Method == http.MethodGet {
		var numberOfBytes, errorMessage = response.Write([]byte(object))
		if numberOfBytes <= 0 {
			log.Printf("Wtf")
		}
		if errorMessage != nil {
			log.Printf("Fuck: %s", errorMessage.Error())
		} else {
			log.Printf("Object: %s", object)
		}
	}
}

//func GetWishList() {
//
//}
