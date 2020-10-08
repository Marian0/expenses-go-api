package expense

//ToExpense -
func ToExpense(expenseDTO DTO) *Expense {
	return &Expense{
		Date:		expenseDTO.Date,
		Details:	expenseDTO.Details,
		Amount:		expenseDTO.Amount,
		User:		expenseDTO.User,
	}
}

//ToExpenseDTO -
func ToExpenseDTO(expense *Expense) DTO {
	return DTO{
		ID:			expense.ID,
		Date:		expense.Date,
		Details:	expense.Details,
		Amount:		expense.Amount,
	}
}

//ToExpenseDTOs -
func ToExpenseDTOs(expenses []Expense) []DTO {
	expenseDTOs := make([]DTO, len(expenses))

	for i, itm := range expenses {
		expenseDTOs[i] = ToExpenseDTO(&itm)
	}

	return expenseDTOs
}
