package handler

import (
	"net/http"
	"encoding/json"
	"gitlab.com/orlmicron/now-test/user"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var u user.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Failed json parse", http.StatusBadRequest)
		return
	}

	user.Create(u)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
