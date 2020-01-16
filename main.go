package main

import (
	"github.com/3pings/clWallApi/route"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/services", route.GetServices).Methods("GET")
	http.ListenAndServe(":5000", router)

}
