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

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}

func Server() {
	r := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8383" // Default port if not specified
	}
	fs := http.FileServer(http.Dir("./web"))
	r.Handle("/web/", http.StripPrefix("/web/", fs))

	r.HandleFunc("/", serveTemplate)
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

//func html(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	if r.Method == "GET" {
//		t, err := template.ParseFiles("web/status_page.html")
//		if err != nil {
//			fmt.Fprintf(w, "parse err")
//			return
//		}
//		t.Execute(w, nil)
//	}
//}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("web", "status_page.html")
	fp := filepath.Join("web", filepath.Clean(r.URL.Path))

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "web", nil)
}
