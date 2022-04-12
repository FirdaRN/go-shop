package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func (server *Server) Initialize(appConfig AppConfig) {
	// func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	fmt.Println("Initializing server..." + appConfig.AppName)
	// var err error

	// if Dbdriver == "mysql" {
	// 	DBURL := "mysql://" + DbUser + ":" + DbPassword + "@" + DbHost + ":" + DbPort + "/" + DbName
	// 	server.DB, err = gorm.Open(Dbdriver, DBURL)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// server.DB.Debug().AutoMigrate(&User{}, &Post{}) //database migration

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	// fmt.Println(http.ListenAndServe(addr, server.Router))
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "go-shop")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "8080")

	server.Initialize(appConfig)
	server.Run(":" + appConfig.AppPort)
}
