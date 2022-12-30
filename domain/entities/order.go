package entities

type OrderEntity struct {
	Price uint
	Title string
}

func NewOrderEntity(price uint, title string) OrderEntity {
	return OrderEntity{Price: price, Title: title}
}
