package main

import (
	"fmt"
	"net/http"

	"github.com/Makinde1034/budget-app/config"
	"github.com/Makinde1034/budget-app/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"

)

func main() {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		AllowedMethods :[]string{"POST", "PUT","GET","DELETE","OPTIONS"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	mux := routes.RegisterRoutes()
	config.Connect()

	err := http.ListenAndServe(":4000",c.Handler(mux))

	fmt.Println(err)

}