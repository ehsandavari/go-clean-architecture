package entities

type OrderEntity struct {
	Id    uint64 `gorm:"primary_key"`
	Price uint
	Title string
}

func (a OrderEntity) TableName() string {
	return "orders"
}
