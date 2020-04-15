package core

import (
	"log"
	"net/http"
)

type PostAudioReq struct {
	IsForcePush bool      `json:"isForcePush"`
	Commands    []Command `json:"commands"`
}

type Command struct {
	Type  int16       `json:"type"`
	Value interface{} `json:"value"`
}

func putCursorNextElement(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json")

	response := "{\"error\": \"Not implemented\"}"
	http.Error(w, response, http.StatusInternalServerError)
}

func putCursorPreviousElement(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json")

	response := "{\"error\": \"Not implemented\"}"
	http.Error(w, response, http.StatusInternalServerError)
}

func putCursorFirstChildElement(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json")

	response := "{\"error\": \"Not implemented\"}"
	http.Error(w, response, http.StatusInternalServerError)
}

func putCursorLastChildElement(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json")

	response := "{\"error\": \"Not implemented\"}"
	http.Error(w, response, http.StatusInternalServerError)
}

func putCursorParentElement(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json")

	response := "{\"error\": \"Not implemented\"}"
	http.Error(w, response, http.StatusInternalServerError)
}
