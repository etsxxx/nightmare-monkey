package nightmare

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (monkey *NightmareMonkey) Play() {
	router := mux.NewRouter()
	r := router.PathPrefix("/api").Subrouter()
	r.Methods("POST").Path("/kill").HandlerFunc(monkey.kill)

	go monkey.Nightmare()

	if err := http.ListenAndServe(fmt.Sprintf(":%d", monkey.ListenPort), router); err != nil {
		log.Fatal(err)
	}
}
