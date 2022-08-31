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
	name := r.URL.Query().Get("name")
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")
	sort := r.URL.Query().Get("sort")

	var sortBool bool

	if sort == "" || sort == "ASC" {
		sortBool = true
	} else if sort == "DESC" {
		sortBool = false
	}

	if page == "" {
		page = "1"
	}

	if size == "" {
		size = "10"
	}

	pageInt, err := strconv.Atoi(page)

	if err != nil {
		helpers.WriteResponse(
			w,
			http.StatusBadRequest,
			"Bad Request",
			nil,
		)
		return

	}

	sizeInt, err := strconv.Atoi(size)

	if err != nil {
		helpers.WriteResponse(
			w,
			http.StatusBadRequest,
			"Bad Request",
			nil,
		)
		return

	}

	result, err := models.GetProducts(pageInt, sizeInt, name, sortBool)

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
