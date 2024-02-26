package main

import (
	"log"
	"net/http"
	"site-done-now/pkg/server"
	"time"
)



func main() {
  address := ":4500";
  mux := http.NewServeMux(); 
  srv := server.NewServer(); 

  mux.HandleFunc("/", srv.HandleIndex);
  mux.HandleFunc("/about", srv.HandleAbout);

  http_server := &http.Server {
    Addr: address,
    Handler: mux, 
    ReadTimeout: 10 * time.Second, 
    WriteTimeout: 10 * time.Second, 
    MaxHeaderBytes: 1 << 20,
  }

  // start the http server on localhost 
  log.Printf("Server listening on localhost:%s...\n", address); 
  log.Fatal(http_server.ListenAndServe()); 
}


