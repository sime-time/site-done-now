package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/a-h/templ"
    "site-done-now/components"
)

func main() {
    const port string = ":4500";
    http.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
 
    index_page := components.Index();
    http.Handle("/", templ.Handler(index_page)); 

    contact_page := components.Contact();
    http.Handle("/contact", templ.Handler(contact_page)); 

    about_page := components.About();
    http.Handle("/about", templ.Handler(about_page));

    fmt.Printf("Server listening on localhost%v... \n", port);
    log.Fatal(http.ListenAndServe(port, nil)); 
}


