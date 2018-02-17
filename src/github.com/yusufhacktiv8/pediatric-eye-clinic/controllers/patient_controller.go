package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yusufhacktiv8/pediatric-eye-clinic/models"
)

// PatientController for patient feature
type PatientController struct {
	DB *sql.DB
}

// CreatePatient create patient
func (a *PatientController) CreatePatient(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&patient); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := patient.Create(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, patient)
}

// FindPatients find patients
func (a *PatientController) FindPatients(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))
	searchText := r.FormValue("searchText")

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	patients, err := models.FindPatients(a.DB, start, count, searchText)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	patientsCount, err := models.CountPatients(a.DB, searchText)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	result := map[string]interface{}{"patients": patients, "count": patientsCount}
	respondWithJSON(w, http.StatusOK, result)
}

// FindPatient to find one patient based on code
func (a *PatientController) FindPatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	patient := models.Patient{Code: code}
	if err := patient.FindOne(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Patient not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, patient)
}

// UpdatePatient update patient
func (a *PatientController) UpdatePatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	var patient models.Patient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&patient); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	patient.Code = code

	if err := patient.Update(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, patient)
}

// DeletePatient delete patient
func (a *PatientController) DeletePatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	patient := models.Patient{Code: code}
	if err := patient.Delete(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
