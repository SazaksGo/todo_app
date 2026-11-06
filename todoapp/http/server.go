package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandlers,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/tasks").Methods("POST").HandlerFunc(s.httpHandlers.HandlerCreateTasks)
	router.Path("/tasks/{title}").Methods("GET").HandlerFunc(s.httpHandlers.HandlerGetTasks)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.httpHandlers.HandlerGetAllTasks)
	router.Path("/tasks").Methods("GET").Queries("completed", "true").HandlerFunc(s.httpHandlers.HandlerGetAllUncompletedTasks)
	router.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandlerCompleteTask)
	router.Path("/tasks/{title}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandlerDeleteTask)

	return http.ListenAndServe(":9091", router)
}
