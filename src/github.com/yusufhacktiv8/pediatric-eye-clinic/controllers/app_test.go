package controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/yusufhacktiv8/pediatric-eye-clinic/models"
)

type AppTest struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func (a *AppTest) Initialize(user, password, dbname string) {
	// connectionString :=
	// 	fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = gorm.Open("postgres", "host=localhost port=5432 user=myyusuf dbname=pec_test sslmode=disable")

	if err != nil {
		log.Fatal(err)
		fmt.Printf("Err: " + err.Error())
	}

	a.DB.AutoMigrate(&models.Role{})

	gin.SetMode(gin.TestMode)
	a.Router = gin.New()
	a.initializeRoutes()
}

func (a *AppTest) initializeRoutes() {
	roleController := RoleController{DB: a.DB}

	v1 := a.Router.Group("/api/roles")
	{
		v1.POST("/", roleController.CreateRole)
		v1.GET("/", roleController.FindRoles)
		v1.PUT("/:id", roleController.UpdateRole)
		v1.DELETE("/:id", roleController.DeleteRole)
	}
}
