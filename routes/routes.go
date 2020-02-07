package routes

import (
	"fmt"
	"net/http"

	"../utils"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	//fs := http.FileServer(http.Dir("./static/"))
	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	return r
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "index.html", struct {
		Title        string
		ASCII        string
		DisplayASCII bool
	}{
		Title:        "Ascii-Art-Web",
		ASCII:        "Nothing yet",
		DisplayASCII: false,
	})
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body := r.PostForm.Get("input-form-textarea")
	fmt.Println(body, []byte(body))
	font := r.PostForm.Get("select-font")
	ReadASCII := utils.GetASCII(font, body)
	utils.ExecuteTemplate(w, "index.html", struct {
		Title        string
		ASCII        string
		DisplayASCII bool
	}{
		Title:        "Ascii-Art-Web",
		ASCII:        ReadASCII,
		DisplayASCII: true,
	})
}
