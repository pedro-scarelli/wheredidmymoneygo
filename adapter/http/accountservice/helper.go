package accountservice

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"strconv"
)

func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}

	return id, nil
}
