package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Rickykn/crud-golang-api.git/helpers"
	"github.com/Rickykn/crud-golang-api.git/models"
	"github.com/gorilla/mux"
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

func GetProductById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		helpers.WriteResponse(
			w,
			http.StatusBadRequest,
			"Bad Request",
			nil,
		)
		return

	}

	result, err := models.GetProductById(id)

	if err == sql.ErrNoRows {
		helpers.WriteResponse(
			w,
			http.StatusNotFound,
			"Data Not Found",
			nil,
		)
		return
	}

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
