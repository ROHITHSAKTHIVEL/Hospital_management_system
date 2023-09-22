package models

type Customer struct {
	FName           string `json:"fname,omitempty"`
	LName           string `json:"lname,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirmpassword,omitempty"`
	//ShippingAddress ShippingAddress `json:"shippingaddress" bson:"shippingaddress"`
	//Is_Valid_User bool `json:"validuser" bson:"validuser"`
}

type Appoitment struct {
	Name string `json:"name,omitempty"`
	// LastName    string `json:"lastname" bson:"lastname"`
	Phonenumber string `json:"phonenumber,omitempty"`
	Purpose     string `json:"purpose,omitempty"`
	Dep         string `json:"dep,omitempty"`
	Email       string `json:"email,omitempty"`
	Date        string `json:"date,omitempty"`
	Time        string `json:"time,omitempty"`
}

type Login struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
