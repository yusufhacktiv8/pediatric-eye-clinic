package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yusufhacktiv8/pediatric-eye-clinic/models"
)

// OccupationController for occupation feature
type OccupationController struct {
	DB *sql.DB
}

// CreateOccupation create occupation
func (a *OccupationController) CreateOccupation(w http.ResponseWriter, r *http.Request) {
	var occupation models.Occupation
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&occupation); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := occupation.Create(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, occupation)
}

// FindOccupations find occupations
func (a *OccupationController) FindOccupations(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	occupations, err := models.FindOccupations(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, occupations)
}

// FindOccupation to find one occupation based on code
func (a *OccupationController) FindOccupation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	occupation := models.Occupation{Code: code}
	if err := occupation.FindOne(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Occupation not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, occupation)
}

// UpdateOccupation update occupation
func (a *OccupationController) UpdateOccupation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	var occupation models.Occupation
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&occupation); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	occupation.Code = code

	if err := occupation.Update(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, occupation)
}

// DeleteOccupation delete occupation
func (a *OccupationController) DeleteOccupation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	occupation := models.Occupation{Code: code}
	if err := occupation.Delete(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
