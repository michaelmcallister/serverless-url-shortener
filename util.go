package main

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type Payload struct {
	Header   int    `json:"reponse"`
	Message  string `json:"message"`
	ShortURL string `json:"shortURL"`
	LongURL  string `json:"longURL"`
}

type RedirObject struct {
	ShortURL string
}

func isValidProtocol(url string) bool {
	validProtocol := regexp.MustCompile(`^https?://`)
	return validProtocol.MatchString(url)
}

func returnResponse(payload *Payload, w http.ResponseWriter, r *http.Request) {
	bytes, _ := json.Marshal(payload)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(payload.Header)
	w.Write(bytes)
	r.Body.Close()
}
