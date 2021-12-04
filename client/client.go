package client

import (
	"encoding/json"
	"fmt"
	"go-distributed-storage/logger"
	"go-distributed-storage/server"
	"io/ioutil"
	"net/http"
	"strings"
)

func Ping(url string) {
	resp, err := http.Get(fmt.Sprint(url, "/server/ping"))
	if err != nil {
		logger.Log(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Log(err.Error())
	}
	logger.Log(string(body))
}

func List(url string) {
	resp, err := http.Get(fmt.Sprint(url, "/data"))
	if err != nil {
		logger.Log(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Log(err.Error())
	}
	logger.Log(string(body))
}

func Get(url string, key string) {
	resp, err := http.Get(fmt.Sprint(url, "/data/get/", key))
	if err != nil {
		logger.Log(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Log(err.Error())
	}
	logger.Log(string(body))
}

func Put(url, key, value string) {
	resp, err := http.Get(fmt.Sprint(url, "/data/put/", key, "/", value))
	if err != nil {
		logger.Log(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Log(err.Error())
	}
	logger.Log(string(body))
}

func Delete(url, key string) {
	resp, err := http.Get(fmt.Sprint(url, "/data/delete/", key))
	if err != nil {
		logger.Log(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Log(err.Error())
	}
	logger.Log(string(body))
}

func Connect(serverUrl, connectUrl string) {
	b, err := json.Marshal(server.ConnectDto{Url: connectUrl})
	resp, err := http.Post(fmt.Sprint(serverUrl, "/server/connect"), "application/json", strings.NewReader(string(b)))
	if err != nil {
		logger.Log(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Log(err.Error())
	}
	logger.Log(string(body))
}
