package routes

import (
	"github.com/Makinde1034/budget-app/controllers"
	"github.com/gorilla/mux"
)

type Method string

const (
	post Method = "POST"
	get Method = "GET"
	patch Method = "PATCH"
	delete Method = "DELETE"
	update Method = "UPDATE"
)

func RegisterRoutes() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/create-budget",controllers.CreateBudget).Methods(string(post))
	router.HandleFunc("/update-budget/{id}",controllers.UpdateBudget).Methods(string(post))
	router.HandleFunc("/budget-activities/{id}",controllers.GetBudgetActivities).Methods(string(get))
	router.HandleFunc("/fetch-budget/{id}",controllers.FetchBudgetId).Methods(string(get))
	router.HandleFunc("/get-budgets",controllers.GetUserBudgets).Methods(string(get))
	// user routes
	router.HandleFunc("/register",controllers.RegisterUser).Methods(string(post))
	router.HandleFunc("/login",controllers.Login).Methods(string(post))

	
	return router

} 