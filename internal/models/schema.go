package models

import (
	"time"

	"github.com/google/uuid"
)

// Branch represents a restaurant location
type Branch struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// Category for auto-categorization
type Category struct {
	ID       uint     `gorm:"primaryKey" json:"id"`
	Name     string   `gorm:"not null" json:"name"`
	Keywords []string `gorm:"type:text[]" json:"keywords"` // Postgres Array
}

// Bill is the "Parent" record that groups multiple expenses
type Bill struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	BranchID     uuid.UUID `gorm:"type:uuid;not null" json:"branch_id"`
	BillingDate  time.Time `gorm:"type:date;not null" json:"billing_date"`
	Source       string    `json:"source"`
	IsSmartInput bool      `gorm:"default:false" json:"is_smart_input"`
	CreatedAt    time.Time `json:"created_at"`

	// Relationship: One Bill has Many Expenses
	// GORM will automatically load these if you use .Preload("Expenses")
	Expenses []Expense `gorm:"foreignKey:BillID;constraint:OnDelete:CASCADE" json:"expenses"`
}

// Expense is the individual "Line Item" inside a bill
type Expense struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	BranchID     uuid.UUID `gorm:"type:uuid;not null" json:"branch_id"`
	BillID       uuid.UUID `gorm:"type:uuid;not null" json:"bill_id"`
	CategoryID   *uint     `json:"category_id"` // Pointer allows null if category is unknown
	EntryDate    time.Time `gorm:"type:date;not null" json:"entry_date"`
	ItemName     string    `gorm:"not null" json:"item_name"`
	Qty          float64   `gorm:"type:decimal" json:"qty"`
	Unit         string    `json:"unit"`
	PricePerUnit float64   `gorm:"type:decimal" json:"price_per_unit"`
	TotalAmount  float64   `gorm:"type:decimal" json:"total_amount"`
	CreatedAt    time.Time `json:"created_at"`
}

// Income represents daily revenue/cash-in
type Income struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	BranchID  uuid.UUID `gorm:"type:uuid;not null" json:"branch_id"`
	EntryDate time.Time `gorm:"type:date;not null" json:"entry_date"`
	Amount    float64   `gorm:"type:decimal" json:"amount"`
	Source    string    `json:"source"` // e.g., "Front Desk", "GrabFood"
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}
