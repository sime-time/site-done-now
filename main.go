package main

import (
	"fmt"
	"log"
	"net/http"
    "site-done-now/server"
)

func main() {
    const address string = ":4500";

    // create server with server-side handler functions
    srv := server.NewServer()  

    http.HandleFunc("/", srv.HandleFilmList)
    http.HandleFunc("/add-film", srv.HandleAddNewFilm)
    http.HandleFunc("/navbar", srv.HandleNavBar)

    fmt.Printf("Server listening on localhost%v... \n", address)
    log.Fatal(http.ListenAndServe(address, nil)); 
}


