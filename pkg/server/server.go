package server 

import "net/http"

// Server is an HTTP server. 
type Server struct {
  
}

// New Server constructor 
func NewServer() *Server {
  return &Server{}; 
}

var indexPage = `
  <!DOCTYPE html>
  <html>
    <body>
      <h1>My First Heading</h1>
      <p>My first paragraph.</p> 
    </body>
  </html>
`;

var aboutPage = `
  <!DOCTYPE html>
  <html>
    <body>
      <h1>About Us</h1>
      <p>My first paragraph about us.</p>
    </body>
  </html>
`;

// HandleIndex handles the index path "/"
func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
  w.Header().Add("content-type", "text/html") 
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(indexPage)) 

}

// HandleAbout handles the path "/about" 
func (s *Server) HandleAbout(w http.ResponseWriter, r *http.Request) {
  w.Header().Add("content-type", "text/html") 
  w.WriteHeader(http.StatusAccepted) 
  w.Write([]byte(aboutPage))
}
