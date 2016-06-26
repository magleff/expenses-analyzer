package cmd

import (
	"fmt"
	"github.com/magleff/gobro/database"
	"github.com/magleff/gobro/features/budget"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

func computeEarnings(budget budget.Budget) float32 {
	earnings := float32(0.00)
	for _, entry := range budget.Expenses {
		if entry.Amount > 0 {
			earnings += entry.Amount
		}
	}
	return earnings
}

func computeExpenses(budget budget.Budget) float32 {
	expenses := float32(0.00)
	for _, entry := range budget.Expenses {
		if entry.Amount <= 0 {
			expenses += entry.Amount
		}
	}
	return expenses
}

func FloatToString(input float32) string {
	return strconv.FormatFloat(float64(input), 'f', 2, 32)
}

func displayBudgetInfos(budget budget.Budget) {
	earnings := computeEarnings(budget)
	expenses := computeExpenses(budget)
	diff := earnings + expenses
	balance := budget.InitialBalance + diff

	fmt.Println("Created on", budget.StartDate)
	fmt.Println("Initial balance", budget.InitialBalance)
	fmt.Println("Total earnings", earnings)
	fmt.Println("Total expenses", expenses)
	fmt.Println("Balance", FloatToString(balance), "("+FloatToString(diff)+")")
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gives the status of the current budget",
	Long:  `Gives the status of the current budget`,
	Run: func(cmd *cobra.Command, args []string) {
		DB := database.NewDatabase()
		controller := budget.NewController(DB)
		currentBudget := controller.CurrentBudget()
		if currentBudget != nil {
			displayBudgetInfos(*currentBudget)
		} else {
			log.Fatal("There is not any active budget.")
		}
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
