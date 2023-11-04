package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebHandler struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type WebServer struct {
	Router        chi.Router
	Handlers      []WebHandler
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      []WebHandler{},
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	s.Handlers = append(s.Handlers, WebHandler{Method: method, Path: path, Handler: handler})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, webHandler := range s.Handlers {
		if webHandler.Method == "GET" {
			s.Router.Get(webHandler.Path, webHandler.Handler)
		}
		if webHandler.Method == "POST" {
			s.Router.Post(webHandler.Path, webHandler.Handler)
		}
		if webHandler.Method == "PUT" {
			s.Router.Put(webHandler.Path, webHandler.Handler)
		}
		if webHandler.Method == "DELETE" {
			s.Router.Delete(webHandler.Path, webHandler.Handler)
		}
		if webHandler.Method == "PATCH" {
			s.Router.Patch(webHandler.Path, webHandler.Handler)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
