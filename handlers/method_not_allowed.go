package handlers

import (
	"net/http"

	"github.com/Rickykn/crud-golang-api.git/helpers"
)

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	helpers.WriteResponse(
		w,
		http.StatusMethodNotAllowed,
		http.StatusText(http.StatusMethodNotAllowed),
		nil,
	)
}
