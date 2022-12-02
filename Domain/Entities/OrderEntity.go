package Entities

type OrderEntity struct {
	Id    uint64 `gorm:"primary_key"`
	Price uint
	Title string
}

func (ne OrderEntity) TableName() string {
	return "orders"
}
