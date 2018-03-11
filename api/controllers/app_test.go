package controllers

import (
	"fmt"
	"log"

	"github.com/appleboy/gofight"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/yusufhacktiv8/pediatric-eye-clinic/models"
	"time"
)

type AppTest struct {
	Router  *gin.Engine
	DB      *gorm.DB
	GoFight *gofight.RequestConfig
}

func (a *AppTest) Initialize(user, password, dbname string) {
	// connectionString :=
	// 	fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	
	connectionParams := "user=docker password=docker dbname=pec sslmode=disable host=db"
	for i := 0; i < 5; i++ {
		a.DB, err = gorm.Open("postgres", connectionParams) // gorm checks Ping on Open
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}
	
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
