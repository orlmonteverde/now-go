package handler

import (
	"net/http"
	"os"
	"fmt"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	secret := os.Getenv("SECRET")
	fmt.Fprintf(w, "Hello world, my secret word is %s", secret)
}
