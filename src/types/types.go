package req_res_types

// Body type response to validate token
type Validate_Token_Body struct {
	Token string `json:"token"`
}

type Costumer struct {
	Full_Name  string `json:"full_name"`
	First_Name string `json:"first_name"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	CPF        string `json:"CPF"`
	IP         string `json:"ip"`
}

type Product struct {
	Product_Id   string `json:"product_id"`
	Product_Name string `json:"product_name"`
}

type Subscription struct {
	ID           string `json:"id"`
	Start_Date   string `json:"start_date"`
	Next_Payment string `json:"next_payment"`
	Status       string `json:"status"`
	Plan         Plan   `json:"plan"`
}

type Plan struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Frenquency  string `json:"frequency"`
	Qty_Changes int16  `json:"qty_charges"`
}

// New Sale type
type KiwifyResponse struct {
	Costumer           Costumer     `json:"Customer"`
	WebHook_Event_type string       `json:"order_approved"`
	Access_url         string       `json:"access_url"`
	Product            Product      `json:"Product"`
	Subscription       Subscription `json:"Subscription"`
	Subscription_ID    string       `json:"subscription_id"`
}

// User in DB
type User struct {
	Data_User *KiwifyResponse `bson:"Data_User"`
	Token     string          `bson:"token"`
}
