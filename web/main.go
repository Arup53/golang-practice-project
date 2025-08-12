package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func handleGet(w http.ResponseWriter , r *http.Request){

	w.Write([]byte("Hello There"))
}

func handlePathName(w http.ResponseWriter, r *http.Request){
    path := chi.URLParam(r, "pathname")
	query := r.URL.Query().Get("q")
	// fmt.Println(r.Body)
	w.Write([]byte("Pathname" + " " + path +" "+ query))
}

func main (){
	r := chi.NewRouter()
    r.Use(middleware.Logger)
	r.Get("/",handleGet)
	r.Get("/dynamicpath/{pathname}",handlePathName )


	http.ListenAndServe(":3000", r)
}