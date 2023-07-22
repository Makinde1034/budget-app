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

	json.NewDecoder(r.Body).Decode(&budgetRequest)
 
	err, resp := model.CreateBudget(&budgetRequest)
                                                                     
	if err != nil {   
		fmt.Println("Erro occured", err)
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

	fmt.Println(resp)
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
	err, resp := model.CreateActivity(&activity)

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

	fmt.Println(resp)

	budget := model.FetchBudgetByID(id)

	fmt.Println(budget.Amount,budget.AmountSpent)


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


	// fmt.Println(activities)
	json.NewEncoder(w).Encode(activities)

}

func GetUserBudgets(w http.ResponseWriter, r *http.Request){

	err, response := model.FetchUserBudgets()

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