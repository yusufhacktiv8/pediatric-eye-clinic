package controllers

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"github.com/yusufhacktiv8/pediatric-eye-clinic/models"
)

func TestFindRoles(t *testing.T) {
	a := GetAppTest()
	r := a.GoFight

	a.DB.Unscoped().Delete(&models.Role{})

	r.GET("/api/roles/").
		Run(a.Router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "{\"message\":\"No count parameter\",\"status\":400}", r.Body.String())
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})

	r.GET("/api/roles/?count=10").
		Run(a.Router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "{\"message\":\"No start parameter\",\"status\":400}", r.Body.String())
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})

	newRole := models.Role{Code: "ADMIN", Name: "Admin"}
	a.DB.Create(&newRole)

	r.GET("/api/roles/?count=10&start=0").
		Run(a.Router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			var objmap map[string]*json.RawMessage
			json.Unmarshal([]byte(r.Body.String()), &objmap)

			var roles []models.Role
			json.Unmarshal(*objmap["data"], &roles)

			role := roles[0]

			assert.Equal(t, "ADMIN", role.Code)
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
