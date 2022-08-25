package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	a "gitlab.com/NijatShahveridev/go_jwt/auth"
	service "gitlab.com/NijatShahveridev/go_jwt/repository"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/signUp", service.SignUp).Methods("POST")
	router.HandleFunc("/signIn", service.SignIn).Methods("POST")
	router.HandleFunc("/admin", a.IsAuthorized(service.AdminIndex)).Methods("GET")
	router.HandleFunc("/user", a.IsAuthorized(service.UserIndex)).Methods("GET")
	router.HandleFunc("/", service.Index).Methods("GET")
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, "+
			"Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, "+
			"Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})

	fmt.Println("Server started at http://localhost:8090")

	err := http.ListenAndServe(":5000",
		handlers.CORS(handlers.AllowedHeaders([]string{
			"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		log.Fatal(err)
	}
}
