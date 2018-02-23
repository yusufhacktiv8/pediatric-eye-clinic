package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/yusufhacktiv8/pediatric-eye-clinic/models"
)

// RoleController for role feature
type RoleController struct {
	DB *gorm.DB
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// FindRoles find roles
func (a *RoleController) FindRoles(w http.ResponseWriter, r *http.Request) {
	// count, _ := strconv.Atoi(r.FormValue("count"))
	// start, _ := strconv.Atoi(r.FormValue("start"))
	// searchText := r.FormValue("searchText")

	var roles []models.Role
	a.DB.Find(&roles)

	result := map[string]interface{}{"roles": roles, "count": 1}

	respondWithJSON(w, http.StatusOK, result)
}

/*
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
	searchText := r.FormValue("searchText")

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	roles, err := models.FindRoles(a.DB, start, count, searchText)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rolesCount, err := models.CountRoles(a.DB, searchText)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	result := map[string]interface{}{"roles": roles, "count": rolesCount}

	respondWithJSON(w, http.StatusOK, result)
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
*/
