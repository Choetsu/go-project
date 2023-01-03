package product

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]Product, error)
	GetByID(id int) (Product, error)
	Create(product Product) (Product, error)
	Update(id int, inputProduct InputProduct) (Product, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB // Database connection
}

func NewProductRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Product, error) { // Get all products from the database
	var products []Product            // Create a slice of products
	err := r.db.Find(&products).Error // Find all products and store them in the products slice
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *repository) GetByID(id int) (Product, error) { // Get a product by its ID
	var product Product                                       // Create a product
	err := r.db.Where(&Product{ID: id}).First(&product).Error // Find the product with the given ID and store it in the product variable
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) Create(product Product) (Product, error) { // Create a new product
	var prod Product
	r.db.Where("name = ?", product.Name).First(&prod)
	if prod.ID != 0 {
		err := errors.New("Ce produit existe déjà !")
		return product, err
	}

	err := r.db.Create(&product).Error // Create the product in the database
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) Update(id int, InputProduct InputProduct) (Product, error) { // Update a product
	var prod Product
	r.db.Where("id = ?", id).First(&prod)
	if prod.ID == 0 {
		err := errors.New("Ce produit n'existe pas !")
		return prod, err
	}

	product, err := r.GetByID(id) // Get the product with the given ID
	if err != nil {
		return product, err // Return the product and the error
	}

	product.Name = InputProduct.Name
	product.Price = InputProduct.Price // Update the product with the new values

	err = r.db.Save(&product).Error // Save the updated product in the database
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) Delete(id int) error {
	product := &Product{ID: id}
	transac := r.db.Delete(product) // Delete the product with the given ID

	if transac.Error != nil {
		return transac.Error
	}
	if transac.RowsAffected == 0 {
		return errors.New("Ce produit n'existe pas !") // Return an error if the product is not found
	}
	return nil
}
