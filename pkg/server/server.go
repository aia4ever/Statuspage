package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"simulator/pkg/result"
)

func Server() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveTemplate)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8383" // Default port if not specified
	}
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

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("web", "status_page.html")
	fp := filepath.Join("web", filepath.Clean(r.URL.Path))

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "status_page", nil)
}
