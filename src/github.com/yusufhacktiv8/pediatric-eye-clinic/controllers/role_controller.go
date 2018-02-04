package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yusufhacktiv8/pediatric-eye-clinic/models"
)

// RoleController for role feature
type RoleController struct {
	DB *sql.DB
}

// CreateRole create role
func (a *RoleController) CreateRole(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&role); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := role.Create(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, role)
}

// FindRoles find roles
func (a *RoleController) FindRoles(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	roles, err := models.FindRoles(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, roles)
}

// FindRole to find one role based on code
func (a *RoleController) FindRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	role := models.Role{Code: code}
	if err := role.FindOne(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Role not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, role)
}

// UpdateRole update role
func (a *RoleController) UpdateRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	var role models.Role
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&role); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	role.Code = code

	if err := role.Update(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, role)
}

// DeleteRole delete role
func (a *RoleController) DeleteRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	role := models.Role{Code: code}
	if err := role.Delete(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
