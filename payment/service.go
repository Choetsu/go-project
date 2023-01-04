package payment

type Service interface {
	GetAll() ([]Payment, error)
	GetByID(id int) (Payment, error)
	Create(payment InputPayment) (Payment, error)
	Update(id int, inputPayment InputPayment) (Payment, error)
	Delete(id int) error
}

type service struct {
	repository Repository // Repository interface
}

func NewPaymentService(r Repository) *service {
	return &service{r}
}

func (s *service) GetAll() ([]Payment, error) { // Get all payments from the database
	payments, err := s.repository.GetAll()
	if err != nil {
		return payments, err
	}
	return payments, nil
}

func (s *service) GetByID(id int) (Payment, error) { // Get a payment by its ID
	payment, err := s.repository.GetByID(id)
	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (s *service) Create(input InputPayment) (Payment, error) { // Create a new payment
	var payment Payment
	payment.ProductID = input.ProductID
	payment.PricePaid = input.PricePaid
	newPayment, err := s.repository.Create(payment)
	if err != nil {
		return payment, err
	}
	return newPayment, nil
}

func (s *service) Update(id int, input InputPayment) (Payment, error) { // Update a payment
	uPayment, err := s.repository.Update(id, input)
	if err != nil {
		return uPayment, err
	}
	return uPayment, nil
}

func (s *service) Delete(id int) error { // Delete a payment
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
