package main

import (
	"fmt"
	"net/http"

	"./routes"
	"./utils"
)

func main() {
	utils.LoadTemplates("templates/*.html")
	r := routes.StandartRouter()
	http.Handle("/", r)
	fmt.Println("Server activated")
	http.ListenAndServe(":8080", nil)
}
