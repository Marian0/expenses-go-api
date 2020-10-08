package expense

import (
	"github.com/marian0/expenses-go-api/components/user"
)

//DTO - Generic Data Transfer Object
type DTO struct {
	ID		uint		`json:"id,string,omitempty"`
	Date	string		`json:"date"`
	Details	string		`json:"details"`
	Amount	int			`json:"amount"`
	User	user.User	`json:"user"`
}

//CreateDTO - Create Data Transfer Object
type CreateDTO struct {
	Date	string			`json:"date" binding:"required"`
	Details	string			`json:"details" binding:"required,lte=100"`
	Amount	int				`json:"amount" binding:"required"`
}

//UpdateDTO - Update Data Transfer Object
type UpdateDTO struct {
	Date	string			`json:"date"`
	Details	string			`json:"details"`
	Amount	int				`json:"amount"`
}
