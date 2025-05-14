package accountservice

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getID(r *http.Request) (string, error) {
	id := mux.Vars(r)["id"]

	return id, nil
}
