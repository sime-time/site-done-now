package server 

import (
    "net/http"
    "html/template"
    "time"
)

type Server struct {
    // write to a database 
}

// constructor 
func NewServer() *Server {
    return &Server{}; 
}

type Film struct {
    Title string 
    Director string 
}

func (s *Server) HandleFilmList(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("views/index.html")) 
    films := map[string][]Film {
        "Films": {
            {Title: "The Godfather", Director: "Francis Ford Coppola"},
            {Title: "Blade Runner", Director: "Ridley Scott"}, 
            {Title: "The Thing", Director: "John Carpenter"},
        },
    }
    tmpl.Execute(w, films) 
}

func (s *Server) HandleAddNewFilm(w http.ResponseWriter, r *http.Request) {
    time.Sleep(1 * time.Second)
    title := r.PostFormValue("title")
    director := r.PostFormValue("director")
    tmpl := template.Must(template.ParseFiles("views/index.html"))
    // reference a block name in our template fragment
    tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
}

