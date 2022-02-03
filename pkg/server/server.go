package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"simulator/pkg/result"
)

func Server() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleConnection)
	s := &http.Server{
		Addr:    "127.0.0.1:8282",
		Handler: r,
	}
	log.Fatal(s.ListenAndServe())

}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	res, err := json.MarshalIndent(result.GetRes(), "", "")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)
}
