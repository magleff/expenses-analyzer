package budget

import (
	"github.com/shopspring/decimal"
	"time"
)

type BudgetInfo struct {
	TotalExpenses          decimal.Decimal
	TotalEarnings          decimal.Decimal
	TotalUncheckedExpenses decimal.Decimal
	InitialBalance         decimal.Decimal
	Difference             decimal.Decimal
	CurrentBalance         decimal.Decimal
	StartDate              time.Time
}
