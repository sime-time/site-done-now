package main

import (
	"fmt"
	"log"
	"net/http"
	"site-done-now/components"
	"github.com/a-h/templ"
)

func main() {
    const address string = ":4500";
    http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./images"))))
 
    index_page := components.Index();
    http.Handle("/", templ.Handler(index_page)); 

    fmt.Printf("Server listening on localhost%v... \n", address);
    log.Fatal(http.ListenAndServe(address, nil)); 
}


