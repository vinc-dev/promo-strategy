package product

// Model struct that represent Product
type Model struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
}

// data master product list that map by product sku
var data = map[string]*Model{
	"120P90": {
		SKU:      "120P90",
		Name:     "Google Home",
		Price:    49.99,
		Quantity: 10,
	},
	"43N23P": {
		SKU:      "43N23P",
		Name:     "MacBook Pro",
		Price:    5399.99,
		Quantity: 5,
	},
	"A304SD": {
		SKU:      "A304SD",
		Name:     "Alexa Speaker",
		Price:    109.50,
		Quantity: 10,
	},
	"234234": {
		SKU:      "234234",
		Name:     "Raspberry Pi B",
		Price:    30,
		Quantity: 2,
	},
}

// dataByName master product list that map by product name
var dataByName = map[string]*Model{
	"Google Home": {
		SKU:      "120P90",
		Name:     "Google Home",
		Price:    49.99,
		Quantity: 10,
	},
	"MacBook Pro": {
		SKU:      "43N23P",
		Name:     "MacBook Pro",
		Price:    5399.99,
		Quantity: 5,
	},
	"Alexa Speaker": {
		SKU:      "A304SD",
		Name:     "Alexa Speaker",
		Price:    109.50,
		Quantity: 10,
	},
	"Raspberry Pi B": {
		SKU:      "234234",
		Name:     "Raspberry Pi B",
		Price:    30,
		Quantity: 2,
	},
}
