package PublishOrder

type SPublishOrderCommand struct {
	Price uint   `json:"price" validate:"required,gte=0,email"`
	Title string `json:"title" validate:"required,gte=0"`
}

func NewSPublishOrderCommand(price uint, title string) SPublishOrderCommand {
	return SPublishOrderCommand{Price: price, Title: title}
}
