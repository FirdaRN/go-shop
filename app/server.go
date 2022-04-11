package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {
	// func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	fmt.Println("Initializing server...")
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

func Run() {
	var server = Server{}
	server.Initialize()
	server.Run(":8080")
}
