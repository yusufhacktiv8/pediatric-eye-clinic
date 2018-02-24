package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yusufhacktiv8/pediatric-eye-clinic/models"
)

// RoleController for role feature
type RoleController struct {
	DB *gorm.DB
}

// FindRoles find roles
func (a *RoleController) FindRoles(c *gin.Context) {
	// count, _ := strconv.Atoi(r.FormValue("count"))
	// start, _ := strconv.Atoi(r.FormValue("start"))
	// searchText := r.FormValue("searchText")

	var roles []models.Role
	a.DB.Find(&roles)

	if len(roles) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No role found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": roles})

}

// CreateRole create role domain model
func (a *RoleController) CreateRole(c *gin.Context) {
	var role models.Role
	c.BindJSON(&role)
	if err := a.DB.Create(&role).Error; err != nil {
		c.AbortWithStatus(400)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"resourceId": role.ID})
}

func (a *RoleController) UpdateRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var role models.Role

	if err := a.DB.Where("id = ?", id).First(&role).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}
	c.BindJSON(&role)
	a.DB.Save(&role)

	c.JSON(http.StatusOK, gin.H{"resourceId": role.ID})
}

func (a *RoleController) DeleteRole(c *gin.Context) {
	id := c.Params.ByName("id")
	var role models.Role
	if err := a.DB.Where("id = ?", id).Delete(&role).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"resourceId": role.ID})
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
