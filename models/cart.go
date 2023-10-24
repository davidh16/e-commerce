package models

type CartItem struct {
	Product     Product `json:"-" gorm:"foreignKey:ProductUuid;references:Uuid"`
	ProductUuid string  `json:"product_uuid"`
	Quantity    int     `json:"quantity"`
	Cart        Cart    `json:"-" gorm:"foreignKey:ProductUuid;references:Uuid"`
	CartUuid    string  `json:"cart_uuid"`
}

type Cart struct {
	Uuid     string `json:"uuid" gorm:"unique;type:uuid; column:uuid;default:uuid_generate_v4()"`
	User     User   `json:"-" gorm:"foreignKey:UserUuid;references:Uuid"`
	UserUuid string `json:"user_uuid"`
	Items    []CartItem
}
