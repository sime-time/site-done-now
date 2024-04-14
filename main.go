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

    services_page := components.Services();
    http.Handle("/services", templ.Handler(services_page)); 

    pricing_page := components.Pricing();
    http.Handle("/pricing", templ.Handler(pricing_page));

    contact_page := components.Contact();
    http.Handle("/contact", templ.Handler(contact_page)); 

    contact_handler := http.HandlerFunc(contactHandler);
    http.Handle("/submit-contact", contact_handler);

    thankyou_page := components.ThankYou();
    http.Handle("/thankyou", templ.Handler(thankyou_page)); 

    fmt.Printf("Server listening on localhost%v... \n", port);
    log.Fatal(http.ListenAndServe(port, nil)); 
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed);
        return 
    }

    // parse form data 
    err := r.ParseForm(); 
    if err != nil {
        http.Error(w, "Error parsing form", http.StatusInternalServerError);
        return
    }

    // get form values 
    name := r.Form.Get("name");
    email := r.Form.Get("email");
    message := r.Form.Get("message"); 

    // validate form data 
    if name == "" || email == "" || message == "" {
        http.Error(w, "Please fill in all fields", http.StatusBadRequest);
        return 
    }

    // send email 

    // success message 
    http.Redirect(w, r, "/thankyou", http.StatusSeeOther); 
}

