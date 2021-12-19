package server

import (
	"encoding/json"
	"fmt"
	"go-distributed-storage/logger"
	"go-distributed-storage/storage"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	store *storage.Storage
	urls  []string
	port  string
	name  string
}

type ConnectDto struct {
	Url string
}

type PongDto struct {
	Name string
}

func New(store *storage.Storage, port string, name string) *Server {
	return &Server{store: store, port: fmt.Sprint(":", port), name: name}
}

func (server *Server) Start() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/data", server.list)
	myRouter.HandleFunc("/data/get/{key}", server.get)
	myRouter.HandleFunc("/data/put/{key}/{value}", server.put)
	myRouter.HandleFunc("/data/delete/{key}", server.put)
	myRouter.HandleFunc("/server/connect", server.connect)
	myRouter.HandleFunc("/server/ping", server.ping)

	go server.pinger()

	logger.Log("Server started")
	logger.Log("Using port", server.port)

	log.Fatal(http.ListenAndServe(server.port, myRouter))
}

func (server *Server) pinger() {
	for {
		for _, url := range server.urls {
			resp, err := http.Get(fmt.Sprint(url, "/server/ping"))
			if err != nil {
				logger.Log(err.Error())
				continue
			}
			server.handlePongResponse(resp)
		}
		time.Sleep(time.Second * 5)
	}
}

func (server *Server) handlePongResponse(r *http.Response) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log(err.Error())
	}
	logger.Log(string(reqBody))
	var pongDto PongDto
	err = json.Unmarshal(reqBody, &pongDto)
	if err != nil {
		logger.Log("Pong response error", err.Error())
	} else {
		logger.Log("Pong:", pongDto.Name)
	}
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

func (server *Server) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	logger.Log("Delete:", key)
	server.store.Delete(key)
	_, _ = fmt.Fprintf(w, "Ok")
}

func (server *Server) connect(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var connectDto ConnectDto
	err := json.Unmarshal(reqBody, &connectDto)
	if err != nil {
		logger.Log("Connect request error", err.Error())
		_, _ = fmt.Fprintf(w, err.Error())
	} else {
		logger.Log("Connect:", connectDto.Url)
		server.urls = append(server.urls, connectDto.Url)
		_, _ = fmt.Fprintf(w, "Ok")
	}
}

func (server *Server) ping(w http.ResponseWriter, r *http.Request) {
	logger.Log("Received ping")
	pong := &PongDto{Name: server.name}
	server.handleError(w, json.NewEncoder(w).Encode(pong))
}

func (server *Server) handleError(w http.ResponseWriter, err error) {
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}
}
