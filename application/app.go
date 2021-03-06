package application

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	model "../model"
	echoPrometheus "github.com/globocom/echo-prometheus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/*ConfigParameters for App*/
type ConfigParameters struct {
	DbName string
	DbUser string
	DbPass string
	DbType string
	DbHost string
	DbPort string
}

/*GoService containing router and database*/
type GoService struct {
	DB     *gorm.DB
	Config *ConfigParameters
}

/*Run app and initialize db connection and http server*/
func (goservice *GoService) Run() {
	//load db connection
	username := goservice.Config.DbUser
	password := goservice.Config.DbPass
	dbName := goservice.Config.DbName
	dbHost := goservice.Config.DbHost
	//dbPort := goservice.Config.DbPort
	dbType := goservice.Config.DbType

	dbURI := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbName)
	//fmt.Println(dbURI)
	conn, err := gorm.Open(dbType, dbURI)
	if err != nil {
		fmt.Print(err)
	}
	//Debug SQL commands
	//conn.LogMode(true)

	goservice.DB = conn
	goservice.DB.Debug().AutoMigrate(&model.Event{})

	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	// Start server
	go func() {
		if err := e.Start(":8000"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(echoPrometheus.MetricsMiddleware())

	// Map Resthandler
	e.GET("/api/v1/event", goservice.GetAll)
	e.POST("/api/v1/event", goservice.PostCreate)
	e.GET("/api/v1/event/:id", goservice.GetByID)
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
