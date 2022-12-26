package PublishOrder

type SPublishOrderCommand struct {
	Price uint   `json:"price" validate:"required"`
	Title string `json:"title" validate:"required"`
}

func NewSPublishOrderCommand(price uint, title string) SPublishOrderCommand {
	return SPublishOrderCommand{Price: price, Title: title}
}
