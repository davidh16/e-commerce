package models

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Uuid            string    `json:"uuid" gorm:"unique;type:uuid; column:uuid;default:uuid_generate_v4()"`
	EmailAddress    string    `json:"email_address"`
	Password        string    `json:"password"`
	ShippingAddress *string   `json:"shipping_address"`
	PaymentInfo     *string   `json:"payment_info" gorm:"embedded"`
	AccountStatus   int       `json:"account_status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	//OrderHistory      []string      `json:"order_history"`
	//Wishlist          []string      `json:"wishlist"`
	//ShoppingCart      []string      `json:"shopping_cart"`
}

func (u *User) Validate() error {
	v := validator.New()
	v.RegisterStructValidationMapRules(ValidationRules, User{})
	err := v.Struct(u)
	if err != nil {
		return err
	}
	return nil
}

var ValidationRules = map[string]string{
	"EmailAddress": "required,email",
	"Password":     "required,min=8",
}

func (u *User) TableName() string {
	return "users"
}

// BeforeCreate - Gorm hook that encrypts password before saving user to database
func (u *User) BeforeCreate(tx *gorm.DB) error {

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}
