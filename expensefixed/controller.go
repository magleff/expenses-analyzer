package expensefixed

import (
	"gopkg.in/mgo.v2"
	"log"
	"strconv"
	"strings"
)

type ExpenseFixedController struct {
	session *mgo.Session
}

func Controller(session *mgo.Session) *ExpenseFixedController {
	return &ExpenseFixedController{session}
}

func (ec ExpenseFixedController) CreateExpenseFixed(amount string, description string) {
	dataStore(ec.session).CreateExpenseFixed(parseAmount(amount), description)
}

func (ec ExpenseFixedController) ListExpensesFixed() []ExpenseFixed {
	return dataStore(ec.session).ListExpensesFixed()
}

// FIXME duplicate code
func parseAmount(amount string) float32 {
	amount = strings.Replace(amount, ",", ".", 1)
	amountFloat, err := strconv.ParseFloat(amount, 32)
	if err != nil {
		log.Fatal(err)
	}
	return float32(amountFloat)
}