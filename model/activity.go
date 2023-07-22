package model

import (
	// "github.com/Makinde1034/budget-app/config"
)

type Activity struct {
	Base
	Title string `json:"title"`
	Amount int `json:"amount"`
	BudgetID string 
}


func CreateActivity(activity *Activity) (error, *Activity)  {
	result := db.Create(&activity)

	if result.Error != nil {
		return result.Error, nil
	}

	return nil, activity

}

func FetchBudgetActivities(budgetID string) (error, [] Activity){
	var activities []Activity
	result := db.Where("budget_id = ?", budgetID).Find(&activities)

	if result.Error != nil {
		return result.Error,nil
	}

	return nil, activities

}