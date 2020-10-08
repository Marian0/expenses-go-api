package expense

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marian0/expenses-go-api/common"
)

//API -
type API struct {
	eService    *Service
}

//NewAPI -
func NewAPI(expenseService *Service) *API {


	return &API{
		eService:	expenseService,
	}
}

//FindAll -
func (eAPI *API) FindAll(c *gin.Context) {
	expenses := eAPI.eService.FindAll()

	common.RespondSuccess(c, http.StatusOK, gin.H{"expenses": ToExpenseDTOs(expenses)})
}

//FindByID -
func (eAPI *API) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	expense := eAPI.eService.FindByID(uint(id))

	common.RespondSuccess(c, http.StatusOK, gin.H{"expense": ToExpenseDTO(&expense)})
}

//Create -
func (eAPI *API) Create(c *gin.Context) {
	//Logged User
	// uuid, _ := c.Get(common.CONST_UUID_IDENTIFIER)

	var expenseDTO CreateDTO

	err := c.ShouldBindJSON(&expenseDTO)
	if err != nil {
		common.RespondError(c, http.StatusBadRequest, err.Error(), "There was a problem trying to create a Expense")
		return
	}

	createdExpense := eAPI.eService.Create(&Expense{
		Date:  expenseDTO.Date,
		Details: expenseDTO.Details,
		Amount: expenseDTO.Amount,
	})

	common.RespondSuccess(c, http.StatusCreated, gin.H{"expense": ToExpenseDTO(createdExpense)})
}

//Update -
func (eAPI *API) Update(c *gin.Context) {
	var expenseDTO UpdateDTO
	err := c.BindJSON(&expenseDTO)
	if err != nil {
		common.RespondError(c, http.StatusBadRequest, err.Error(), "There was a problem trying to update a Expense")
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	expense := eAPI.eService.FindByID(uint(id))
	if expense == (Expense{}) {
		common.RespondError(c, http.StatusNotFound, "Expense not found "+c.Param("id"), "Expense not found")
		return
	}

	expense.Date = expenseDTO.Date
	expense.Details = expenseDTO.Details
	expense.Amount = expenseDTO.Amount

	eAPI.eService.Save(expense)

	common.RespondSuccess(c, http.StatusOK, gin.H{"expense": ToExpenseDTO(&expense)})
}

//Delete -
func (eAPI *API) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	expense := eAPI.eService.FindByID(uint(id))
	if expense == (Expense{}) {
		common.RespondError(c, http.StatusNotFound, "Expense not found "+c.Param("id"), "Expense not found")
		return
	}

	eAPI.eService.Delete(expense)

	common.RespondSuccess(c, http.StatusOK, gin.H{"expense": ToExpenseDTO(&expense)})
}
