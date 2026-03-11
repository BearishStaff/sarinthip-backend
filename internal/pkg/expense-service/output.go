package expensesvc

import (
	"time"

	"github.com/google/uuid"
)

type ExpenseOutput struct {
	ID           uuid.UUID `json:"id"`
	BranchID     uuid.UUID `json:"branch_id"`
	BillID       uuid.UUID `json:"bill_id"`
	Category     string    `json:"category"`
	EntryDate    time.Time `json:"entry_date"`
	ItemName     string    `json:"item_name"`
	Quantity     float64   `json:"qty"`
	Unit         string    `json:"unit"`
	PricePerUnit float64   `json:"price_per_unit"`
	TotalAmount  float64   `json:"total_amount"`
	CreatedAt    time.Time `json:"created_at"`
}
