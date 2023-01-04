package payment

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]Payment, error)
	GetByID(id int) (Payment, error)
	Create(payment Payment) (Payment, error)
	Update(id int, inputPayment InputPayment) (Payment, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB // Database connection
}

func NewPaymentRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Payment, error) { // Get all payments from the database
	var payments []Payment            // Create a slice of payments
	err := r.db.Find(&payments).Error // Find all payments and store them in the payments slice
	if err != nil {
		return payments, err
	}
	return payments, nil
}

func (r *repository) GetByID(id int) (Payment, error) { // Get a payment by its ID
	var payment Payment                                       // Create a payment
	err := r.db.Where(&Payment{ID: id}).First(&payment).Error // Find the payment with the given ID and store it in the payment variable
	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *repository) Create(payment Payment) (Payment, error) { // Create a new payment
	err := r.db.Create(&payment).Error // Create the payment in the database
	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *repository) Update(id int, inputPayment InputPayment) (Payment, error) { // Update a payment
	payment, err := r.GetByID(id) // Get the payment with the given ID
	if err != nil {
		return payment, err // Return the payment and the error
	}

	payment.ProductID = inputPayment.ProductID
	payment.PricePaid = inputPayment.PricePaid // Update the payment with the new values
	r.db.Save(&payment)                        // Save the updated payment in the database

	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *repository) Delete(id int) error { // Delete a payment
	payment := &Payment{ID: id} // Get the payment with the given ID
	transac := r.db.Delete(payment)

	if transac.Error != nil {
		return transac.Error
	}
	if transac.RowsAffected == 0 {
		return errors.New("payment non trouvé")
	}
	return nil
}
