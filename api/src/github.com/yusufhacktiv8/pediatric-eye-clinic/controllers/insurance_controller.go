package controllers

/*
import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yusufhacktiv8/pediatric-eye-clinic/models"
)

// InsuranceController for insurance feature
type InsuranceController struct {
	DB *sql.DB
}

// CreateInsurance create insurance
func (a *InsuranceController) CreateInsurance(w http.ResponseWriter, r *http.Request) {
	var insurance models.Insurance
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&insurance); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := insurance.Create(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, insurance)
}

// FindInsurances find insurances
func (a *InsuranceController) FindInsurances(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))
	searchText := r.FormValue("searchText")

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	insurances, err := models.FindInsurances(a.DB, start, count, searchText)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	insurancesCount, err := models.CountInsurances(a.DB, searchText)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	result := map[string]interface{}{"insurances": insurances, "count": insurancesCount}
	respondWithJSON(w, http.StatusOK, result)
}

func (a *InsuranceController) FindAllInsurances(w http.ResponseWriter, r *http.Request) {
	insurances, err := models.FindAllInsurances(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, insurances)
}

// FindInsurance to find one insurance based on code
func (a *InsuranceController) FindInsurance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	insurance := models.Insurance{Code: code}
	if err := insurance.FindOne(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Insurance not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, insurance)
}

// UpdateInsurance update insurance
func (a *InsuranceController) UpdateInsurance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	var insurance models.Insurance
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&insurance); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	insurance.Code = code

	if err := insurance.Update(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, insurance)
}

// DeleteInsurance delete insurance
func (a *InsuranceController) DeleteInsurance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	insurance := models.Insurance{Code: code}
	if err := insurance.Delete(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
*/
