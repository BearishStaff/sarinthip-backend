package expensesvc

import (
	"time"

	"github.com/google/uuid"
)

type IExpenseService interface {
	GetExpenses(branchID uuid.UUID) ([]ExpenseOutput, error)
}

type ExpenseService struct{}

func NewExpenseService() *ExpenseService {
	return &ExpenseService{}
}

func (service *ExpenseService) GetExpenses(branchID uuid.UUID) ([]ExpenseOutput, error) {
	if branchID.String() == "9ec944d0-3120-43d3-82e8-bb7adfad6622" {
		return []ExpenseOutput{
			{
				ID:           uuid.Must(uuid.NewUUID()),
				BranchID:     branchID,
				BillID:       uuid.Must(uuid.NewUUID()),
				Category:     "หมู",
				EntryDate:    time.Now(),
				ItemName:     "ตับหมู",
				Quantity:     3,
				Unit:         "กก.",
				PricePerUnit: 120,
				TotalAmount:  360,
				CreatedAt:    time.Now(),
			},
			{
				ID:           uuid.Must(uuid.NewUUID()),
				BranchID:     branchID,
				BillID:       uuid.Must(uuid.NewUUID()),
				Category:     "หมู",
				EntryDate:    time.Now(),
				ItemName:     "สามชั้น",
				Quantity:     10,
				Unit:         "กก.",
				PricePerUnit: 250,
				TotalAmount:  2500,
				CreatedAt:    time.Now(),
			},
			{
				ID:           uuid.Must(uuid.NewUUID()),
				BranchID:     branchID,
				BillID:       uuid.Must(uuid.NewUUID()),
				Category:     "เนื้อ",
				EntryDate:    time.Now(),
				ItemName:     "เนื้อสับ",
				Quantity:     5,
				Unit:         "กก.",
				PricePerUnit: 260,
				TotalAmount:  1300,
				CreatedAt:    time.Now(),
			},
		}, nil
	}
	return []ExpenseOutput{
		{
			ID:           uuid.Must(uuid.NewUUID()),
			BranchID:     branchID,
			BillID:       uuid.Must(uuid.NewUUID()),
			Category:     "เส้น",
			EntryDate:    time.Now(),
			ItemName:     "เส้นหมี่",
			Quantity:     10,
			Unit:         "แพ็ค",
			PricePerUnit: 100,
			TotalAmount:  1000,
			CreatedAt:    time.Now(),
		},
		{
			ID:           uuid.Must(uuid.NewUUID()),
			BranchID:     branchID,
			BillID:       uuid.Must(uuid.NewUUID()),
			Category:     "ผัก",
			EntryDate:    time.Now(),
			ItemName:     "ผักบุ้ง",
			Quantity:     5,
			Unit:         "กก.",
			PricePerUnit: 50,
			TotalAmount:  250,
			CreatedAt:    time.Now(),
		},
	}, nil
}
