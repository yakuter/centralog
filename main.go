package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/yakuter/centralog/pkg/database"
	"github.com/yakuter/centralog/pkg/entry"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func main() {
	db, err := database.New("centralog.db")
	if err != nil {
		log.Fatal(err)
	}

	app := &App{
		Router: mux.NewRouter(),
		DB:     db,
	}

	app.Router.HandleFunc("/", app.list).Methods("GET")
	app.Router.HandleFunc("/log", app.log).Methods("POST")
	app.Router.HandleFunc("/health", app.health).Methods("GET")

	fmt.Println("Server is running on port 8090")
	http.ListenAndServe(":8090", app.Router)
}

func (a *App) list(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	entries, err := database.List(a.DB)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(entries); err != nil {
		fmt.Println("Error: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) log(w http.ResponseWriter, req *http.Request) {
	var e entry.Entry
	if err := json.NewDecoder(req.Body).Decode(&e); err != nil {
		fmt.Println("Error: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.Insert(a.DB, &e); err != nil {
		fmt.Println("Error: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
