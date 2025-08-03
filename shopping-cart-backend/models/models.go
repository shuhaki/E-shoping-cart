package models

import(
	"time"
	go.mongodb.org/mongo-driver/bson/primitive
)  
	

	typr User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	First_name  string             `bson:"username" json:"username"`
	Last_name   string             `bson:"last_name" json:"last_name"`
	Password  string             `bson:"password" json:"password"`
	Email     string             `bson:"email" json:"email"`
	Phone     string             `bson:"phone" json:"phone"`
	Token     string             `bson:"token" json:"token"`
	RefreshToken string             `bson:"refresh_token" json:"refresh_token"`

	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	User_ID   primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Usercart []CartItem        `bson:"user_cart" json:"user_cart"`
	Address_Details []AddressDetails `bson:"address_details" json:"address_details"`
	order_status string `bson:"order_status" json:"order_status"`
}
type Product struct {
	
	ProductID primitive.ObjectID `bson:"product_id" json:"product_id"`
	ProductName string             `bson:"product_name" json:"product_name"`
	Price float64            `bson:"price" json:"price"`
	Rating float64           `bson:"rating" json:"rating"`
	ImageURL string            `bson:"image_url" json:"image_url"`
}
type ProductUser struct {
	ProductID primitive.ObjectID `bson:"product_id" json:"product_id"`
	ProductName string             `bson:"product_name" json:"product_name"`
	Price float64            `bson:"price" json:"price"`
	Rating float64           `bson:"rating" json:"rating"`
	ImageURL string            `bson:"image_url" json:"image_url"`




}
type Address struct {
	Address_id primitive.ObjectID `bson:"address_id" json:"address_id"`
	House_number string `bson:"house_number" json:"house_number"`
	Street string `bson:"street" json:"street"`
	City string `bson:"city" json:"city"`
	Pin_code string `bson:"pin_code" json:"pin_code"`

}

type Order struct {
	OrderID primitive.ObjectID `bson:"order_id" json:"order_id"`
	Order_Cart []CartItem `bson:"order_cart" json:"order_cart"`
	Ordered_At time.Time `bson:"ordered_at" json:"ordered_at"`
	Price float64 `bson:"price" json:"price"`
	Discount float64 `bson:"discount" json:"discount"`
	Payment_Method string `bson:"payment_method" json:"payment_method"`

}

type Payment struct {
	Digital bool
	COD bool	
}

