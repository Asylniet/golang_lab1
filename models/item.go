package models

import (
	"fmt"
)

type Item struct {
	ID     int
	Name   string
	Price  int
	Rating float32
}

// constructor
func NewItem(name string, price int, rating float32) *Item {
	return &Item{
		Name:   name,
		Price:  price,
		Rating: rating,
	}
}

// toString
func (item *Item) Print(id int) {
	fmt.Printf("%v. %s with rating %v costs %v\n", id, item.Name, item.Rating, item.Price)
}

// Setters
func (item *Item) SetName(name string) {
	item.Name = name
}

func (item *Item) SetRating(rating float32) {
	item.Rating = rating
}

func (item *Item) SetPrice(price int) {
	item.Price = price
}
