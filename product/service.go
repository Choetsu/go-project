package product

type Service interface {
	GetAll() ([]Product, error)
	GetByID(id int) (Product, error)
	Create(product InputProduct) (Product, error)
	Update(id int, inputProduct InputProduct) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository // Repository interface
}

func NewProductService(r Repository) *service {
	return &service{r}
}

func (s *service) GetAll() ([]Product, error) { // Get all products from the database
	products, err := s.repository.GetAll()
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *service) GetByID(id int) (Product, error) { // Get a product by its ID
	product, err := s.repository.GetByID(id)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *service) Create(input InputProduct) (Product, error) { // Create a new product
	var product Product
	product.Name = input.Name
	product.Price = input.Price
	newProduct, err := s.repository.Create(product)
	if err != nil {
		return product, err
	}
	return newProduct, nil
}

func (s *service) Update(id int, input InputProduct) (Product, error) { // Update a product
	uProduct, err := s.repository.Update(id, input)
	if err != nil {
		return uProduct, err
	}
	return uProduct, nil
}

func (s *service) Delete(id int) error { // Delete a product
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
