package main

import (
	"net/http"
	"log"
	"github.com/gorilla/sessions"
	"github.com/mapu77/AD-Labs/6-go-webapp/routes"
)

func main() {

	sessions.NewCookieStore([]byte("travelAgencyCookie"))
	router := routes.GetRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}
