package routes

import (
	"fmt"
	"net/http"

	"../utils"
)

func StandartRouter() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/", StdIndexHandler)

	//fs := http.FileServer(http.Dir("./static/"))
	//http.Handle("/css/", http.StripPrefix("/css/", fs))
	return r
}

func StdIndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound) //404 error
		return
	}
	switch r.Method {
	case http.MethodGet:
		indexGetHandler(w, r)
	case http.MethodPost:
		indexPostHandler(w, r)
	default:
	}
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
	if !utils.IsAlphanumerical(body) {
		w.WriteHeader(http.StatusBadRequest)
		//w.Write([]byte("Status Bad Request"))
		errorHandler(w, r, http.StatusBadRequest) //400 error
		return
	}
	font := r.PostForm.Get("select-font")
	ReadASCII, err := utils.GetASCII(font, body)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		//w.Write([]byte("Status Internal Server Error"))
		errorHandler(w, r, http.StatusInternalServerError) //500 error
		return
	}
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

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	if status == http.StatusNotFound {
		http.ServeFile(w, r, "./templates/404.html")
	} else if status == http.StatusBadRequest {
		http.ServeFile(w, r, "./templates/400.html")
	} else {
		http.ServeFile(w, r, "./templates/500.html")
	}
}

// func NewRouter() *mux.Router {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", indexGetHandler).Methods("GET")
// 	r.HandleFunc("/", indexPostHandler).Methods("POST")
// 	//fs := http.FileServer(http.Dir("./static/"))
// 	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
// 	return r
// }
