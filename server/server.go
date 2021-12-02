package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-distributed-storage/logger"
	"go-distributed-storage/storage"
	"log"
	"net/http"
)

type Server struct {
	store *storage.Storage
	urls  []string
}

func New(store *storage.Storage) *Server {
	return &Server{store: store}
}

func (server *Server) Start() {
	server.store.Put("Kalle", "Kula")
	server.store.Put("Bertil", "Svensson")

	var port = ":8080"
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/data", server.list)
	myRouter.HandleFunc("/data/{key}", server.get)
	myRouter.HandleFunc("/data/put/{key}/{value}", server.put)
	myRouter.HandleFunc("/server/connect/{url}", server.connect)
	myRouter.HandleFunc("/server/ping", server.ping)

	logger.Log("Server started")
	logger.Log("Using port", port)

	log.Fatal(http.ListenAndServe(port, myRouter))
}

func (server *Server) list(w http.ResponseWriter, r *http.Request) {
	logger.Log("Handling request", r.RequestURI)
	server.handleError(w, json.NewEncoder(w).Encode(server.store.List()))
}

func (server *Server) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	logger.Log("Fetching", key)
	data, exists := server.store.Get(key)
	if exists {
		server.handleError(w, json.NewEncoder(w).Encode(data))
	} else {
		server.handleError(w, json.NewEncoder(w).Encode(nil))
	}
}

func (server *Server) put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]
	logger.Log("Put:", key, ":", value)
	server.store.Put(key, value)
	_, _ = fmt.Fprintf(w, "Ok")
}

func (server *Server) connect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := vars["url"]
	logger.Log("Connect:", url)
	server.urls = append(server.urls, url)
	_, _ = fmt.Fprintf(w, "Ok")
}

func (server *Server) ping(w http.ResponseWriter, r *http.Request) {
	logger.Log("Received ping")
	_, _ = fmt.Fprintf(w, "pong")
}

func (server *Server) handleError(w http.ResponseWriter, err error) {
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}
}
