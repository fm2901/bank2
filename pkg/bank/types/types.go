package types

// Amount of money in minimum units (diram, kopeyka, cent...)
type Money int64

// Category of payment
type Category string

// Structure of payment
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
