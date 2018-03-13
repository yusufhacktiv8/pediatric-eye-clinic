package main

import (
	"fmt"
	"log"

	"time"

	"github.com/yusufhacktiv8/pediatric-eye-clinic/controllers"
	"github.com/yusufhacktiv8/pediatric-eye-clinic/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func (a *App) Initialize(user, password, dbname string) {
	// connectionString :=
	// 	fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	// a.DB, err = gorm.Open("postgres", "host=db port=5432 user=docker password=docker dbname=pec sslmode=disable")

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

	a.Router = gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	a.Router.Use(cors.New(config))
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	roleController := controllers.RoleController{DB: a.DB}

	v1 := a.Router.Group("/api/roles")
	{
		v1.POST("/", roleController.CreateRole)
		v1.GET("/", roleController.FindRoles)
		v1.PUT("/:id", roleController.UpdateRole)
		v1.DELETE("/:id", roleController.DeleteRole)
	}
}

// func (a *App) Run(addr string) {
// 	c := cors.New(cors.Options{
// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
// 		AllowedHeaders:   []string{"*"},
// 		AllowCredentials: true,
// 	})
// 	handler := c.Handler(a.Router)
// 	log.Fatal(http.ListenAndServe(addr, handler))
// }
