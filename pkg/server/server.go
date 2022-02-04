package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"simulator/pkg/result"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func Server() {
	r := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8383"
	}
	r.Handle("/", http.FileServer(http.Dir("./static")))
	r.HandleFunc("/api", handleConnection)
	s := &http.Server{
		Addr:    ":" + port,
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
