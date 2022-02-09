package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"simulator/pkg/emulator"
	"simulator/pkg/result"
	"time"
)

func main() {
	r := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8383"
	}
	ticker := time.NewTicker(30 * time.Second)
	var res result.ResultT
	go func() {
		for range ticker.C {
			emulator.Shuffle()
			res = result.GetRes()
		}
	}()
	r.HandleFunc("/api", handleConnection(&res))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	s := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	log.Fatal(s.ListenAndServe())

}

func handleConnection(res *result.ResultT) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		json, err := json.MarshalIndent(res, "", "")
		if err != nil {
			log.Fatal(err)
		}
		w.Write(json)
	}
}
