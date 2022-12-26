package SubscribeOrder

type SSubscribeOrderCommand struct {
	Price uint   `json:"price" validate:"required"`
	Title string `json:"title" validate:"required"`
}

func NewSSubscribeOrderCommand(price uint, title string) SSubscribeOrderCommand {
	return SSubscribeOrderCommand{Price: price, Title: title}
}
