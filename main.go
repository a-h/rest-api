package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PostRequest struct {
	Name string `json:"name"`
}

type PostResponse struct {
	Message string `json:"message"`
}

func NewHomeHandler(greeter func(name string) string) HomeHandler {
	return HomeHandler{
		Greeter: greeter,
	}
}

type HomeHandler struct {
	// Put your database clients here.
	Greeter func(name string) string
}

func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var postRequest PostRequest
	err := json.NewDecoder(r.Body).Decode(&postRequest)
	if err != nil {
		http.Error(w, "failed to decode body", http.StatusBadRequest)
		return
	}
	resp := PostResponse{
		Message: h.Greeter(postRequest.Name),
	}
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func main() {
	greeter := func(name string) string {
		return fmt.Sprintf("你好，%s", name)
	}
	http.Handle("/", NewHomeHandler(greeter))
	http.ListenAndServe("localhost:8000", nil)
}
