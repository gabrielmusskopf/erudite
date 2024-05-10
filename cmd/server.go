package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	URL string
}

func (s Server) Post(path string, body any) (*http.Response, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", s.URL+path, bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 10 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return res, nil
}

func (s Server) Get(path string) (*http.Response, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	r, err := http.Get(s.URL + path)
	if err != nil {
		return nil, err
	}

	return r, nil
}
