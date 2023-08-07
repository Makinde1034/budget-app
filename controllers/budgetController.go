package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Makinde1034/budget-app/model"
	"github.com/gorilla/mux"
	// "github.com/Makinde1034/budget-app/helpers"
)

func CreateBudget(w http.ResponseWriter, r *http.Request){          

	budgetRequest := model.Budget{}
	userId,_ := r.Context().Value("id").(string)

	json.NewDecoder(r.Body).Decode(&budgetRequest)

	budget := model.Budget{
		Name : budgetRequest.Name,
		Description : budgetRequest.Description,
		StartDate : budgetRequest.StartDate,
		EndDate : budgetRequest.EndDate,
		Amount : budgetRequest.Amount,
		Owner: userId,
		Color: budgetRequest.Color,
	}
 
	err, _ := model.CreateBudget(&budget)
                                                                     
	if err != nil {   
		fmt.Println("Error occured", err)
		errResponse := struct{
			Msg string `json:"msg"` 
		}{ 
			"error creating budget",
		} 
		json.NewEncoder(w).Encode(errResponse)
	}

	response := struct{
		Msg string `json:"msg"`
	}{
		"Budget successfully created",
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}

func UpdateBudget(w http.ResponseWriter, r *http.Request){          
	activityRequest := model.Activity{}
	
	params := mux.Vars(r)
	id := params["id"]
	json.NewDecoder(r.Body).Decode(&activityRequest)

	activity := model.Activity{
		Amount: activityRequest.Amount,
		BudgetID: id,
		Title: activityRequest.Title,
	}

	// create activity related to budget
	err, _ := model.CreateActivity(&activity)

	if err != nil {
		fmt.Println("Erro occured", err)
		errResponse := struct{
			Msg string `json:"msg"`
		}{
			"error updating budget",
		}
		json.NewEncoder(w).Encode(errResponse)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}


	budget := model.FetchBudgetByID(id)
	// update budget
	
	model.UpdateBudget(activityRequest.Amount + budget.AmountSpent,id)

	response := struct{
		Msg string `json:"msg"`
	}{
		"Budget successfully updated",
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	
}

func GetBudgetActivities(w http.ResponseWriter, r *http.Request) { 
	
	params := mux.Vars(r)
	id := params["id"]

	err,activities := model.FetchBudgetActivities(id)

	if err != nil {
		fmt.Println(err,"errrrrr")
		errResponse := struct{
			Msg string `json:"msg"`
		}{
			"error updating budget",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errResponse)
		
		return
	}


	
	json.NewEncoder(w).Encode(activities)

}

func GetUserBudgets(w http.ResponseWriter, r *http.Request){
	
	userId,_ := r.Context().Value("id").(string)

	err, response := model.FetchUserBudgets(userId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrMsg{"Failed to get user budgets."})
	}

	json.NewEncoder(w).Encode(response)
}

func FetchBudgetId( w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id := params["id"]

	response := model.FetchBudgetByID(id)
	json.NewEncoder(w).Encode(response)
}