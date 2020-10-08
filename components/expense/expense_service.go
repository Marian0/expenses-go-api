package expense

//Service -
type Service struct {
	eRepo	*Repository
}

//NewService -
func NewService(eRepo *Repository) *Service {
	return &Service{eRepo: eRepo}
}

//FindAll -
func (eService *Service) FindAll() []Expense {
	return eService.eRepo.FindAll()
}

//FindByID -
func (eService *Service) FindByID(id uint) Expense {
	return eService.eRepo.FindByID(id)
}

//Create -
func (eService *Service) Create(expense *Expense) *Expense {
	return eService.eRepo.Create(expense)
}

//Save -
func (eService *Service) Save(expense Expense) Expense {
	eService.eRepo.Save(expense)

	return expense
}

//Delete -
func (eService *Service) Delete(expense Expense) {
	eService.eRepo.Delete(expense)
}
