package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"math/rand"
	_ "html/template"
	_ "path/filepath"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"log"
)

//func Index(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "Welcome!")
//}

func Index(w http.ResponseWriter, r *http.Request) {
	ren := render.New()
	ren.HTML(w,http.StatusOK, "index2", "World")
	//templ, err := template.ParseFiles(filepath.Join("templates","index.html"))
	//if err != nil {
	//	panic(err)
	//}
	//templ.Execute(w, nil)

	//t.once.Do(func() {
	//	t.templ =  template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))
	//})
	//t.templ.Execute(w, nil)
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func dumbShow(w http.ResponseWriter, r *http.Request) {
	rating := rand.Intn(100)
	vars := mux.Vars(r)
	name := vars["name"]
	dumbs := Dumb{
		Name: name,
		Dumb: rating,
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(dumbs); err != nil {
		panic(err)
	}
}

func dumbShow2(w http.ResponseWriter, r *http.Request){
	ren := render.New()
	rating := rand.Intn(100)
	vars := mux.Vars(r)
	name := vars["name"]

	dumbs := Dumb{
		Name: name,
		Dumb: rating,
	}
	ren.JSON(w,http.StatusOK, dumbs)
}
