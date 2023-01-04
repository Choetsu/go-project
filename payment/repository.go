package payment

import (
	"errors"
	"go-project/product"

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

	for i, payment := range payments {
		r.db.Where("id = ?", payment.ProductID).First(&payments[i].Product) // Find the product with the given ID and store it in the product variable
	}

	if err != nil {
		return payments, err
	}
	return payments, nil
}

func (r *repository) GetByID(id int) (Payment, error) { // Get a payment by its ID
	var payment Payment                                       // Create a payment
	err := r.db.Where(&Payment{ID: id}).First(&payment).Error // Find the payment with the given ID and store it in the payment variable

	r.db.Where(&product.Product{ID: payment.ProductID}).First(&payment.Product) // Find the product with the given ID and store it in the product variable
	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *repository) Create(payment Payment) (Payment, error) { // Create a new payment
	r.db.Where("id = ?", payment.ProductID).First(&payment.Product)

	if payment.Product.ID == 0 {
		return payment, errors.New("Produit non trouvé")
	}

	var payments Payment
	var product product.Product
	payments.ProductID = product.ID
	payments.PricePaid = product.Price
	payments.Product = product

	err := r.db.Create(&payment).Error // Create the payment in the database
	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *repository) Update(id int, inputPayment InputPayment) (Payment, error) { // Update a payment
	var product product.Product
	r.db.Where("id = ?", inputPayment.ProductID).First(&product)

	payment, err := r.GetByID(id) // Get the payment with the given ID
	if err != nil {
		return payment, err // Return the payment and the error
	}

	payment.ProductID = product.ID             // Update the payment with the new values
	payment.PricePaid = inputPayment.PricePaid // Update the payment with the new values
	payment.Product = product                  // Update the payment with the new values
	err = r.db.Save(&payment).Error            // Save the updated payment in the database

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
