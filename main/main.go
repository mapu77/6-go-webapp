package main

import (
	"net/http"
	"log"
	"github.com/gorilla/sessions"
	"github.com/mapu77/AD-Labs/6-go-webapp/routes"
	"os"
)

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}
	sessions.NewCookieStore([]byte("travelAgencyCookie"))
	router := routes.GetRouter()

	log.Fatal(http.ListenAndServe(":"+port, router))

}
