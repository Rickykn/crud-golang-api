package handlers

import (
	"net/http"

	"github.com/Rickykn/crud-golang-api.git/helpers"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	helpers.WriteResponse(
		w,
		http.StatusNotFound,
		"Oops nothing to be seen here...",
		nil,
	)
}
