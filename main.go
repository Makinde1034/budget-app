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
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	mux := routes.RegisterRoutes()
	config.Connect()

	err := http.ListenAndServe(":4000",c.Handler(mux))

	fmt.Println(err)

}