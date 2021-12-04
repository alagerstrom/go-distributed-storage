package client

import (
	"fmt"
	"go-distributed-storage/logger"
	"io/ioutil"
	"net/http"
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
