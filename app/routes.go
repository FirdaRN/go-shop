package app

import "github.com/FirdaRN/go-shop/app/controllers"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
	// server.Router.HandleFunc("/users", server.createUser).Methods("POST")
	// server.Router.HandleFunc("/users", server.getUsers).Methods("GET")
	// server.Router.HandleFunc("/users/{id}", server.getUser).Methods("GET")
	// server.Router.HandleFunc("/users/{id}", server.updateUser).Methods("PUT")
	// server.Router.HandleFunc("/users/{id}", server.deleteUser).Methods("DELETE")
}