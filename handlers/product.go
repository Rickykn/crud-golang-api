package handlers

import (
	"net/http"

	"github.com/Rickykn/crud-golang-api.git/helpers"
	"github.com/Rickykn/crud-golang-api.git/models"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	result, err := models.GetProducts()

	if err != nil {
		helpers.WriteResponse(
			w,
			http.StatusInternalServerError,
			"Server Error",
			nil,
		)
		return
	}
	helpers.WriteResponse(
		w,
		http.StatusOK,
		"Get Data Success",
		result,
	)
}
