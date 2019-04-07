package handler

import (
	"net/http"
	"encoding/json"
	"gitlab.com/orlmicron/now-test/user"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	users := user.GetAll()
	j, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Failed json parse", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
