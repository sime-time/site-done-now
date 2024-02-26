package main 

import (
  "net/http"
  "log"
  "time" 
)

func handler(w http.ResponseWriter, r *http.Request) {
  var indexPage = `
    <!DOCTYPE html>
    <html>
      <body>
        <h1>My First Heading</h1>
        <p>My first paragraph.</p> 
      </body>
    </html>
  `;
  w.Header().Add("content-type", "text/html") 
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(indexPage)) 
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
  var aboutPage = `
    <!DOCTYPE html>
    <html>
      <body>
        <h1>About Us</h1>
        <p>My first paragraph about us.</p>
      </body>
    </html>
  `;
  w.Header().Add("content-type", "text/html") 
  w.WriteHeader(http.StatusAccepted) 
  w.Write([]byte(aboutPage))
}

func main() {
  const address string = ":4500";
  mux := http.NewServeMux(); 

  mux.HandleFunc("/", handler);
  mux.HandleFunc("/about", aboutHandler);

  server := &http.Server {
    Addr: address,
    Handler: mux, 
    ReadTimeout: 10 * time.Second, 
    WriteTimeout: 10 * time.Second, 
    MaxHeaderBytes: 1 << 20,
  }

  // start the http server on localhost 
  log.Printf("Server listening on localhost:%s...\n", address); 
  log.Fatal(server.ListenAndServe()); 

}


