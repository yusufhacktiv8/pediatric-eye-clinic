package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yusufhacktiv8/pediatric-eye-clinic/models"
)

// MedicalRecordController for medicalRecord feature
type MedicalRecordController struct {
	DB *sql.DB
}

// CreateMedicalRecord create medicalRecord
func (a *MedicalRecordController) CreateMedicalRecord(w http.ResponseWriter, r *http.Request) {
	var medicalRecord models.MedicalRecord
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&medicalRecord); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := medicalRecord.Create(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, medicalRecord)
}

// FindMedicalRecords find medicalRecords
func (a *MedicalRecordController) FindMedicalRecords(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))
	searchText := r.FormValue("searchText")

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	medicalRecords, err := models.FindMedicalRecords(a.DB, start, count, searchText)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	medicalRecordsCount, err := models.CountMedicalRecords(a.DB, searchText)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	result := map[string]interface{}{"medicalRecords": medicalRecords, "count": medicalRecordsCount}

	respondWithJSON(w, http.StatusOK, result)
}

// FindMedicalRecord to find one medicalRecord based on code
func (a *MedicalRecordController) FindMedicalRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	medicalRecord := models.MedicalRecord{Code: code}
	if err := medicalRecord.FindOne(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "MedicalRecord not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, medicalRecord)
}

// UpdateMedicalRecord update medicalRecord
func (a *MedicalRecordController) UpdateMedicalRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	var medicalRecord models.MedicalRecord
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&medicalRecord); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	medicalRecord.Code = code

	if err := medicalRecord.Update(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, medicalRecord)
}

// DeleteMedicalRecord delete medicalRecord
func (a *MedicalRecordController) DeleteMedicalRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	medicalRecord := models.MedicalRecord{Code: code}
	if err := medicalRecord.Delete(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
