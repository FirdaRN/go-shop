package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FirdaRN/go-shop/database/seeders"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"gorm.io/driver/mysql"
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

type DBConfig struct {
	DBHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DBPort     string
	DBDriver   string
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Initializing server..." + appConfig.AppName)

	server.initializeRoutes()
}

func (server *Server) initializeDB(dbConfig DBConfig) {
	fmt.Println("Initializing database...")
	var err error

	if dbConfig.DBDriver == "mysql" {
		server.DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.DbUser, dbConfig.DbPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DbName)), &gorm.Config{})
	}

	if err != nil {
		panic(err)
	}
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	// fmt.Println(http.ListenAndServe(addr, server.Router))
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) dbMigrate() {
	for _, model := range RegisterModels() {
		err := server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database initialized")
}

func (server *Server) initCommands(config AppConfig, dbConfig DBConfig) {
	server.initializeDB(dbConfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeders.DBSeed(server.DB)
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
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
	var dbConfig = DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "go-shop")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "8080")

	dbConfig.DBHost = getEnv("DB_HOST", "127.0.0.1")
	dbConfig.DbUser = getEnv("DB_USER", "root")
	dbConfig.DbPassword = getEnv("DB_PASSWORD", "")
	dbConfig.DbName = getEnv("DB_NAME", "goshop")
	dbConfig.DBPort = getEnv("DB_PORT", "3306")
	dbConfig.DBDriver = getEnv("DB_DRIVER", "mysql")

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.initCommands(appConfig, dbConfig)
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)
	}
}
