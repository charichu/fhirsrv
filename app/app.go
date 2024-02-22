package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type App struct {
	Router *http.ServeMux
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	a.Router = http.NewServeMux()
	a.InitializeRoutes()
}

func (a *App) Run(addr string) {
	fmt.Println("FHIR has been lit on the server")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("GET /path/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "got path\n")
	})
	a.Router.HandleFunc("/task/{id}/", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "handling task with id=%v\n", id)
	})
}
