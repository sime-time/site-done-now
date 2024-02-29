package server

import (
	"io"
	"net/http"
    "log"
    "encoding/json"
)

// Server is an HTTP server.
type Server struct {
    // map to simulate writing to a database server 
    users map[string]userInfo // key -> username 
}

// New Server constructor 
func NewServer() *Server {
  return &Server{
        users: make(map[string]userInfo),
    }; 
}


var indexPage = `
  <!DOCTYPE html>
  <html>
    <body>
      <h1>User Database</h1>
      <p>Welcome to the user database.</p> 
    </body>
  </html>
`;

// user represents the JSON value that's sent as a response to a request. 
type user struct {
    Name string `json:"name"`
    Email string `json:"email"` 
    Age int `json:"age"`
}


type userInfo struct {
    email string 
    age int 
}


// HandleIndex handles the index path "/"
func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
  w.Header().Add("content-type", "text/html") 
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(indexPage)) 

}

// HandleCreateUser handles the path "/users/create" 
// Create -> Post or Put 
func (s *Server) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost, http.MethodPut:
        // Check that the input is JSON 
        contentType := r.Header.Get("content-type") 
        if contentType != "application/json" {
            w.WriteHeader(http.StatusUnsupportedMediaType) 
            return 
        }

        body, err := io.ReadAll(r.Body) 
        if err != nil {
            log.Printf("Could not read request body: %v", err) 
            w.WriteHeader(http.StatusInternalServerError) // HTTP 500 
            return
        }

        defer r.Body.Close() 

        // Unmarshall the body 
        var u user 
        err2 := json.Unmarshal(body, &u) 
        if err2 != nil {
            log.Printf("Could not unmarshal request body: %v", err2)
            w.WriteHeader(http.StatusBadRequest) // HTTP 400 
            return 
        }
       
        log.Printf("Create User: %v", u.Name) 
        s.users[u.Name] = userInfo{
            email: u.Email,
            age: u.Age, 
        }
        
    default: 
        w.WriteHeader(http.StatusMethodNotAllowed) // HTTP 415 
    }
}

// HandleReadUsers handles the "/users" request for fetching a user
func (s *Server) HandleReadUsers(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    // fetch the name from the query string 
    case http.MethodGet: 
        name := r.URL.Query().Get("name")
        u, ok := s.users[name] 
        if !ok {
            w.WriteHeader(http.StatusNotFound) 
            return
        }
        ret := user {
            Name: name,
            Email: u.email,
            Age: u.age, 
        }
        msg, err := json.Marshal(ret) 
        if err != nil {
            log.Printf("Could not marshal request: %v", err) 
            w.WriteHeader(http.StatusInternalServerError) // HTTP 500 
            return
        } 
        w.Header().Add("Content-Type", "application/json")
        w.Write(msg)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed) // HTTP 415 
    } 
}
