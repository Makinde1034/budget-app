package model

import (
	"fmt"
	"gorm.io/gorm"
)

var db *gorm.DB


type Budget struct {
	Base
	Name string `json:"name"`
	Description string `json:"description"`
	Amount int `gorm:"default:0" json:"amount"`
	StartDate string `json:"startDate"`
	EndDate string `json:"endDate"`
	DaysLeft string `json:"daysLeft"`
	AmountSpent int `json:"amountSpent"`
	Color string `json:"color"`
	Owner string `json:"owner"`

}

func CreateBudget(budget *Budget) (error, *Budget)  {
	result := db.Create(&budget)

	if result.Error != nil {
		return result.Error, nil
	}

	return nil, budget

}

func UpdateBudget(amountSpent int,budgetId string) {
	// var budget Budget
	db.Model(Budget{}).Where("id = ?", budgetId).Updates(Budget{AmountSpent: amountSpent})
}

func FetchBudgetByID(id string) Budget {
	var budget Budget
	fmt.Println(id)
	db.First(&budget, "ID = ?", id)
	return budget
}

func FetchUserBudgets(ownerId  string) (error, []Budget) {
	var budgets []Budget;

	result := db.Where("Owner = ?", ownerId).Find(&budgets)

	if result.Error != nil {
		return result.Error, nil
	}

	return nil, budgets

}