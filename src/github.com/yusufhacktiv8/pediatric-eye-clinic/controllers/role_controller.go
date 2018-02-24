package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yusufhacktiv8/pediatric-eye-clinic/models"
)

// RoleController for role feature
type RoleController struct {
	DB *gorm.DB
}

// FindRoles find roles within range defined by start and count query parameters
// Use the searchText query parameter to filter code or name (incasesensitive)
func (a *RoleController) FindRoles(c *gin.Context) {
	countStr, ok := c.GetQuery("count")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "No count parameter"})
		return
	}

	startStr, ok := c.GetQuery("start")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "No start parameter"})
		return
	}

	count, _ := strconv.Atoi(countStr)
	start, _ := strconv.Atoi(startStr)
	searchText, _ := c.GetQuery("searchText")
	searchText = strings.ToLower(searchText)

	var roles []models.Role
	a.DB.Where("lower(code) LIKE ?", "%"+searchText+"%").Or("lower(name) LIKE ?", "%"+searchText+"%").Offset(start).Limit(count).Find(&roles)

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
	if err := a.DB.Where("id = ?", id).First(&role).Delete(&role).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"resourceId": role.ID})
}
