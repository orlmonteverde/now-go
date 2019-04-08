package handler

import (
	"net/http"
	"encoding/json"
	"strconv"
	"gitlab.com/orlmicron/now-test/user"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Failed id parse", http.StatusBadRequest)
		return
	}

	u := user.GetOne(id)
	j, err := json.Marshal(u)
	if err != nil {
		http.Error(w, "Failed json parse", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
