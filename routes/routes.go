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
	// Route untuk register dan login 
	r.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", controllers.LoginUser).Methods("POST")

	r.HandleFunc("/users", controllers.GetUsers).Subrouter()
	r.Use(middleware.ValidateToken)

	// Route CRUD untuk users dengan autentikasi
	r.HandleFunc("/", controllers.GetUser).Methods("GET")
	r.HandleFunc("/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/{id}", controllers.DeleteUser).Methods("delete")
	r.Use(middleware.LoggingMiddleware)

	// Menjalankan server dan menangani error jika gagal
	log.Println("server run on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server Failed to start", err)
	}
}
