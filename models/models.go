package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	Name            string  `json:"name" bson:"name"`
	Phone_No        int     `json:"phonenumber" bson:"phonenumber"`
	Age             int     `json:"age" bson:"age"`
	Password        string  `json:"password" bson:"password"`
	ConfirmPassword string  `json:"confirmpassword" bson:"confirmpassword"`
	//ShippingAddress ShippingAddress `json:"shippingaddress" bson:"shippingaddress"`
	Is_Valid_User bool `json:"validuser" bson:"validuser"`
}

type ShippingAddress struct {
	FirstName   string `json:"firstname" bson:"firstname"`
	LastName    string `json:"lastname" bson:"lastname"`
	House_No    string `json:"houseno" bson:"houseno"`
	Street_Name string `json:"streetname" bson:"streetname"`
	City        string `json:"city" bson:"city"`
	Pincode     int64  `json:"pincode" bson:"pincode"`
}

type Orders struct {
	Item_id    string  `json:"itemid" bson:"itemid"`
	Item_Name  string  `json:"itemname" bson:"itemname"`
	Quantity   int64   `json:"quantity" bson:"quantity"`
	Total_Cost float64 `json:"totalcost" bson:"totalcost"`
}
type Customer_Response struct {
	Id               primitive.ObjectID `json:"_id" bson:"_id"`
	Phone_No         int64              `json:"phonenumber" bson:"phonenumber"`
	Age              int                `json:"age" bson:"age"`
	Balance          float64            `json:"balance" bson:"balance"`
	Shipping_Address ShippingAddress    `json:"shippingadderss" bson:"shipingaddress"`
	Is_Valid_User    bool               `json:"validuser" bson:"validuser"`
	Orders           []Orders           `json:"orders" bson:"orders"`
}
