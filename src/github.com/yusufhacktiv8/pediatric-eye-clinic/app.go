package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/yusufhacktiv8/pediatric-eye-clinic/controllers"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		// fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
		fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
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

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := getProducts(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (a *App) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	p := product{ID: id}
	if err := p.getProduct(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Product not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) createProduct(w http.ResponseWriter, r *http.Request) {
	var p product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.createProduct(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func (a *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var p product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	p.ID = id

	if err := p.updateProduct(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	p := product{ID: id}
	if err := p.deleteProduct(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
	a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.getProduct).Methods("GET")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")

	diseaseController := controllers.DiseaseController{DB: a.DB}
	patientController := controllers.PatientController{DB: a.DB}
	medicalRecordController := controllers.MedicalRecordController{DB: a.DB}
	occupationController := controllers.OccupationController{DB: a.DB}
	insuranceController := controllers.InsuranceController{DB: a.DB}
	roleController := controllers.RoleController{DB: a.DB}
	userController := controllers.UserController{DB: a.DB}
	a.Router.HandleFunc("/diseases", diseaseController.CreateDisease).Methods("POST")
	a.Router.HandleFunc("/diseases", diseaseController.FindDiseases).Methods("GET")
	a.Router.HandleFunc("/diseases/{code:\\w+}", diseaseController.FindDisease).Methods("GET")
	a.Router.HandleFunc("/diseases/{code:\\w+}", diseaseController.UpdateDisease).Methods("PUT")
	a.Router.HandleFunc("/diseases/{code:\\w+}", diseaseController.DeleteDisease).Methods("DELETE")

	a.Router.HandleFunc("/patients", patientController.CreatePatient).Methods("POST")
	a.Router.HandleFunc("/patients", patientController.FindPatients).Methods("GET")
	a.Router.HandleFunc("/patients_all", patientController.FindAllPatients).Methods("GET")
	a.Router.HandleFunc("/patients/{code:\\w+}", patientController.FindPatient).Methods("GET")
	a.Router.HandleFunc("/patients/{code:\\w+}", patientController.UpdatePatient).Methods("PUT")
	a.Router.HandleFunc("/patients/{code:\\w+}", patientController.DeletePatient).Methods("DELETE")

	a.Router.HandleFunc("/medicalrecords", medicalRecordController.CreateMedicalRecord).Methods("POST")
	a.Router.HandleFunc("/medicalrecords", medicalRecordController.FindMedicalRecords).Methods("GET")
	a.Router.HandleFunc("/medicalrecords/{code:\\w+}", medicalRecordController.FindMedicalRecord).Methods("GET")
	a.Router.HandleFunc("/medicalrecords/{code:\\w+}", medicalRecordController.UpdateMedicalRecord).Methods("PUT")
	a.Router.HandleFunc("/medicalrecords/{code:\\w+}", medicalRecordController.DeleteMedicalRecord).Methods("DELETE")

	a.Router.HandleFunc("/occupations", occupationController.CreateOccupation).Methods("POST")
	a.Router.HandleFunc("/occupations", occupationController.FindOccupations).Methods("GET")
	a.Router.HandleFunc("/occupations_all", occupationController.FindAllOccupations).Methods("GET")
	a.Router.HandleFunc("/occupations/{code:\\w+}", occupationController.FindOccupation).Methods("GET")
	a.Router.HandleFunc("/occupations/{code:\\w+}", occupationController.UpdateOccupation).Methods("PUT")
	a.Router.HandleFunc("/occupations/{code:\\w+}", occupationController.DeleteOccupation).Methods("DELETE")

	a.Router.HandleFunc("/insurances", insuranceController.CreateInsurance).Methods("POST")
	a.Router.HandleFunc("/insurances", insuranceController.FindInsurances).Methods("GET")
	a.Router.HandleFunc("/insurances_all", insuranceController.FindAllInsurances).Methods("GET")
	a.Router.HandleFunc("/insurances/{code:\\w+}", insuranceController.FindInsurance).Methods("GET")
	a.Router.HandleFunc("/insurances/{code:\\w+}", insuranceController.UpdateInsurance).Methods("PUT")
	a.Router.HandleFunc("/insurances/{code:\\w+}", insuranceController.DeleteInsurance).Methods("DELETE")

	a.Router.HandleFunc("/roles", roleController.CreateRole).Methods("POST")
	a.Router.HandleFunc("/roles", roleController.FindRoles).Methods("GET")
	a.Router.HandleFunc("/roles/{code:\\w+}", roleController.FindRole).Methods("GET")
	a.Router.HandleFunc("/roles/{code:\\w+}", roleController.UpdateRole).Methods("PUT")
	a.Router.HandleFunc("/roles/{code:\\w+}", roleController.DeleteRole).Methods("DELETE")

	a.Router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	a.Router.HandleFunc("/users", userController.FindUsers).Methods("GET")
	a.Router.HandleFunc("/users/{email:\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3}}", userController.UpdateUser).Methods("PUT")
	a.Router.HandleFunc("/users/{email:\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3}}", userController.DeleteUser).Methods("DELETE")
}

func (a *App) Run(addr string) {
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})
	handler := c.Handler(a.Router)
	log.Fatal(http.ListenAndServe(addr, handler))
}
