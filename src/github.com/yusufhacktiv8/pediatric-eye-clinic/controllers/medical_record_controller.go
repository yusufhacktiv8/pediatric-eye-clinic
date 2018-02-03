package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yusufhacktiv8/pediatric-eye-clinic/models"
)

// MedicalRecordController for patient feature
type MedicalRecordController struct {
	DB *sql.DB
}

// CreateMedicalRecord create patient
func (a *MedicalRecordController) CreateMedicalRecord(w http.ResponseWriter, r *http.Request) {
	var patient models.MedicalRecord
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

// FindMedicalRecords find patients
func (a *MedicalRecordController) FindMedicalRecords(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	patients, err := models.FindMedicalRecords(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, patients)
}

// FindMedicalRecord to find one patient based on code
func (a *MedicalRecordController) FindMedicalRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	patient := models.MedicalRecord{Code: code}
	if err := patient.FindOne(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "MedicalRecord not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, patient)
}

// UpdateMedicalRecord update patient
func (a *MedicalRecordController) UpdateMedicalRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	var patient models.MedicalRecord
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

// DeleteMedicalRecord delete patient
func (a *MedicalRecordController) DeleteMedicalRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	patient := models.MedicalRecord{Code: code}
	if err := patient.Delete(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
