package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.lancs.ac.uk/library/auditor/pkg/events"
	"log"
	"net/http"
)

type App interface {
	Run(addr string)
}

type app struct {
	Router *mux.Router
	Repo events.Repository
}

func (a *app) Run(addr string) {
	fmt.Printf("API Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *app) initializeRoutes() {
	a.Router.HandleFunc("/audit_events", a.getAuditEvents).Methods("GET")
}

func (a *app) getAuditEvents(w http.ResponseWriter, r *http.Request) {

	e, err := a.Repo.GetAll()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, e)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func NewAPIService(repo events.Repository) App {
	a := &app{
		Repo: repo,
		Router: mux.NewRouter(),
	}

	a.initializeRoutes()

	return a
}
