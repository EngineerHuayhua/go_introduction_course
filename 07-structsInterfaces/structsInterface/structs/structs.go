package structs

type Product struct {
	ID    uint     `json:"id"`
	Name  string   `json:"name"`
	Type  Type     `json:"type"`
	Count uint16   `json:"count"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

type Type struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

// TotalPrice calculates the total price of the product based on its count and price.
// It returns the total price as a float64 value.
func (p Product) TotalPrice() float64 {
	return float64(p.Count) * p.Price
}

// SetName sets the Name field of the Product struct.
// It takes a string parameter and updates the Name field accordingly.
func (p *Product) SetName(name string) {
	p.Name = name
}

// AddTags appends one or more tags to the Product's Tags slice.
// It takes a variadic parameter, allowing multiple tags to be added at once.
func (p *Product) AddTags(tags ...string) {
	p.Tags = append(p.Tags, tags...)
}

// Car represents a collection of products owned by a user.
// It contains an ID, a UserID to associate it with a user, and a slice
type Car struct {
	ID       uint      `json:"id"`
	UserID   uint      `json:"user_id"`
	Products []Product `json:"products"`
}

// AddProduct adds one or more products to the Car's Products slice.
// It appends the provided products to the existing slice.
func (c *Car) AddProduct(product ...Product) {
	c.Products = append(c.Products, product...)
}

// Total calculates the total price of all products in the Car.
// It iterates through the Products slice and sums up the total price of each product.
func (c Car) Total() float64 {
	total := 0.0
	for _, product := range c.Products {
		// For each product, it calls the TotalPrice method to get the total price
		// and adds it to the total.
		total += product.TotalPrice()
	}
	return total
}

// NewCar creates a new Car instance with the specified userID.
// It initializes the Car with the given userID and returns it.
func NewCar(userID uint) Car {
	return Car{UserID: userID}
}
