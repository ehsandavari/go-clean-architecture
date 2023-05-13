package createOrder

type SCreateOrderCommand struct {
	UserId uint
	Title  string
	Price  uint
}

func NewSCreateOrderCommand(userId uint, title string, price uint) SCreateOrderCommand {
	return SCreateOrderCommand{
		UserId: userId,
		Title:  title,
		Price:  price,
	}
}
