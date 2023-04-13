package models

type Basket struct {
	ID     int
	UserId int
	ItemId int
}

// constructor
func NewBasket(userId, itemId int) *Basket {
	return &Basket{
		UserId: userId,
		ItemId: itemId,
	}
}
