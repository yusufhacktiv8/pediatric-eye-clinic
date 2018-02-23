package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/yusufhacktiv8/pediatric-eye-clinic/controllers"
	"github.com/yusufhacktiv8/pediatric-eye-clinic/models"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(user, password, dbname string) {
	// connectionString :=
	// 	// fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	// 	fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	// a.DB, err = sql.Open("postgres", connectionString)
	a.DB, err = gorm.Open("postgres", "host=localhost port=5432 user=myyusuf dbname=pec sslmode=disable")

	if err != nil {
		log.Fatal(err)
		fmt.Printf("Err: " + err.Error())
	}

	a.DB.AutoMigrate(&models.Role{})

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

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {

	jwtSecretKey := "jshbdgh54gs9jdbx543GnhY67"

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(jwtSecretKey), nil
				})
				if error != nil {
					respondWithError(w, http.StatusBadRequest, "Error in request")
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					respondWithError(w, http.StatusBadRequest, "Invalid authorization token")
				}
			}
		} else {
			respondWithError(w, http.StatusBadRequest, "An authorization header is required")
		}
	})
}

func (a *App) initializeRoutes() {
	// loginController := controllers.LoginController{DB: a.DB}
	// diseaseController := controllers.DiseaseController{DB: a.DB}
	// patientController := controllers.PatientController{DB: a.DB}
	// medicalRecordController := controllers.MedicalRecordController{DB: a.DB}
	// occupationController := controllers.OccupationController{DB: a.DB}
	// insuranceController := controllers.InsuranceController{DB: a.DB}
	roleController := controllers.RoleController{DB: a.DB}
	// userController := controllers.UserController{DB: a.DB}

	// a.Router.HandleFunc("/authenticate", loginController.Authenticate).Methods("POST")
	//
	// a.Router.HandleFunc("/diseases", diseaseController.CreateDisease).Methods("POST")
	// a.Router.HandleFunc("/diseases", diseaseController.FindDiseases).Methods("GET")
	// a.Router.HandleFunc("/diseases/{code:\\w+}", diseaseController.FindDisease).Methods("GET")
	// a.Router.HandleFunc("/diseases/{code:\\w+}", diseaseController.UpdateDisease).Methods("PUT")
	// a.Router.HandleFunc("/diseases/{code:\\w+}", diseaseController.DeleteDisease).Methods("DELETE")
	//
	// a.Router.HandleFunc("/patients", patientController.CreatePatient).Methods("POST")
	// a.Router.HandleFunc("/patients", ValidateMiddleware(patientController.FindPatients)).Methods("GET")
	// a.Router.HandleFunc("/patients_all", patientController.FindAllPatients).Methods("GET")
	// a.Router.HandleFunc("/patients/{code:\\w+}", patientController.FindPatient).Methods("GET")
	// a.Router.HandleFunc("/patients/{code:\\w+}", patientController.UpdatePatient).Methods("PUT")
	// a.Router.HandleFunc("/patients/{code:\\w+}", patientController.DeletePatient).Methods("DELETE")
	//
	// a.Router.HandleFunc("/medicalrecords", medicalRecordController.CreateMedicalRecord).Methods("POST")
	// a.Router.HandleFunc("/medicalrecords", medicalRecordController.FindMedicalRecords).Methods("GET")
	// a.Router.HandleFunc("/medicalrecords/{code:\\w+}", medicalRecordController.FindMedicalRecord).Methods("GET")
	// a.Router.HandleFunc("/medicalrecords/{code:\\w+}", medicalRecordController.UpdateMedicalRecord).Methods("PUT")
	// a.Router.HandleFunc("/medicalrecords/{code:\\w+}", medicalRecordController.DeleteMedicalRecord).Methods("DELETE")
	//
	// a.Router.HandleFunc("/occupations", occupationController.CreateOccupation).Methods("POST")
	// a.Router.HandleFunc("/occupations", occupationController.FindOccupations).Methods("GET")
	// a.Router.HandleFunc("/occupations_all", occupationController.FindAllOccupations).Methods("GET")
	// a.Router.HandleFunc("/occupations/{code:\\w+}", occupationController.FindOccupation).Methods("GET")
	// a.Router.HandleFunc("/occupations/{code:\\w+}", occupationController.UpdateOccupation).Methods("PUT")
	// a.Router.HandleFunc("/occupations/{code:\\w+}", occupationController.DeleteOccupation).Methods("DELETE")
	//
	// a.Router.HandleFunc("/insurances", insuranceController.CreateInsurance).Methods("POST")
	// a.Router.HandleFunc("/insurances", insuranceController.FindInsurances).Methods("GET")
	// a.Router.HandleFunc("/insurances_all", insuranceController.FindAllInsurances).Methods("GET")
	// a.Router.HandleFunc("/insurances/{code:\\w+}", insuranceController.FindInsurance).Methods("GET")
	// a.Router.HandleFunc("/insurances/{code:\\w+}", insuranceController.UpdateInsurance).Methods("PUT")
	// a.Router.HandleFunc("/insurances/{code:\\w+}", insuranceController.DeleteInsurance).Methods("DELETE")

	// a.Router.HandleFunc("/roles", roleController.CreateRole).Methods("POST")
	a.Router.HandleFunc("/roles", roleController.FindRoles).Methods("GET")
	// a.Router.HandleFunc("/roles/{code:\\w+}", roleController.FindRole).Methods("GET")
	// a.Router.HandleFunc("/roles/{code:\\w+}", roleController.UpdateRole).Methods("PUT")
	// a.Router.HandleFunc("/roles/{code:\\w+}", roleController.DeleteRole).Methods("DELETE")
	//
	// a.Router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	// a.Router.HandleFunc("/users", userController.FindUsers).Methods("GET")
	// a.Router.HandleFunc("/users/{email:\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3}}", userController.UpdateUser).Methods("PUT")
	// a.Router.HandleFunc("/users/{email:\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3}}", userController.DeleteUser).Methods("DELETE")
}

func (a *App) Run(addr string) {
	c := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(a.Router)
	log.Fatal(http.ListenAndServe(addr, handler))
}
