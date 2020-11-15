package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Coaster struct {
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	ID           string `json:"id"`
	InPark       string `json:"in_park"`
	Height       int    `json:"height"`
}

type CoasterHandlers struct {
	store map[string]Coaster
}


func newCoasterHandlers() *CoasterHandlers {
	h := CoasterHandlers{
		store: make(map[string]Coaster),
	}
	for i:=0; i < 10; i++ {
		h.store[strconv.Itoa(i)] = Coaster {
			Name: strconv.Itoa(i),
		}
	}
	return &h
}


func (h *CoasterHandlers) coasters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	case http.MethodPost:
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed."))
		return
	}
}


func (h *CoasterHandlers) get(w http.ResponseWriter, r *http.Request) {
	coasters := make([]Coaster, len(h.store))
	i := 0
	for _, coaster := range h.store {
		coasters[i] = coaster
		i++
	}
	jsonBytes, err := json.Marshal(coasters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}


func (h *CoasterHandlers) post(w http.ResponseWriter, r *http.Request) {
}


func main() {
	coasterHandlers := newCoasterHandlers()
	http.HandleFunc("/coasters", coasterHandlers.coasters)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
