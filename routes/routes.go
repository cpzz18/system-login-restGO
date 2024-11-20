package routes

import (
	"myapi/controllers"
	"myapi/middleware"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

)

func Route() {
	r := mux.NewRouter()
	r.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", controllers.LoginUser).Methods("POST")
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("delete")
	r.Use(middleware.LoggingMiddleware)

	log.Println("server run")
	http.ListenAndServe(":8000", r)
}
