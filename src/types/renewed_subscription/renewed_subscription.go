package renewedsubscription

type Renewed_subscription struct {
	Order_id              string             `json:"order_id"`
	Order_ref             string             `json:"order_ref"`
	Order_status          string             `json:"order_status"`
	Product_type          string             `json:"product_type"`
	Payment_method        string             `json:"payment_method"`
	Store_id              string             `json:"store_id"`
	Payment_merchant_Id   string             `json:"payment_merchant_id"`
	Installments          int8               `json:"installments"`
	Card_type             string             `json:"card_type"`
	Card_last4digits      string             `json:"card_last4digits"`
	Card_rejection_reason string             `json:"card_rejection_reason"`
	Boleto_URL            string             `json:"boleto_URL"`
	Boleto_barcode        string             `json:"boleto_barcode"`
	Boleto_expiry_date    string             `json:"boleto_expiry_date"`
	Pix_code              string             `json:"pix_code"`
	Pix_expiration        string             `json:"pix_expiration"`
	Sale_type             string             `json:"sale_type"`
	Created_at            string             `json:"created_at"`
	Updated_at            string             `json:"updated_at"`
	Approved_date         string             `json:"approved_date"`
	Refunded_at           string             `json:"refunded_at"`
	Webhook_event_type    string             `json:"webhook_event_type"`
	Product               Product            `json:"Product"`
	Costumer              Costumer           `json:"Customer"`
	Commissions           Commissions        `json:"Commissions"`
	TrackingParameters    TrackingParameters `json:"TrackingParameters"`
	Subscription          Subscription       `json:"Subscription"`
	Subscription_id       string             `json:"subscription_id"`
}

type Product struct {
	Product_id   string `json:"product_id"`
	Product_name string `json:"product_name"`
}

type Costumer struct {
	Full_name  string `json:"full_name"`
	First_name string `json:"first_name"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	CPF        string `json:"CPF"`
	Ip         string `json:"ip"`
}

type Commissions struct {
	Charge_amount          string                `json:"charge_amount"`
	Product_base_price     string                `json:"product_base_price"`
	Kiwify_fee             string                `json:"kiwify_fee"`
	Commissioned_stores    []Commissioned_stores `json:"commissioned_stores"`
	Currency               string                `json:"currency"`
	My_commission          string                `json:"my_commission"`
	Funds_status           string                `json:"funds_status"`
	Estimated_deposit_date string                `json:"estimated_deposit_date"`
	Deposit_date           string                `json:"deposit_date"`
}

type Commissioned_stores struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Custom_name string `json:"custom_name"`
	Email       string `json:"email"`
	Value       string `json:"value"`
}

type TrackingParameters struct {
	Src          string `json:"src"`
	Sck          string `json:"sck"`
	Utm_source   string `json:"utm_source"`
	Utm_medium   string `json:"utm_medium"`
	Utm_campaign string `json:"utm_campaign"`
	Utm_content  string `json:"utm_content"`
	Utm_term     string `json:"utm_term"`
}

type Subscription struct {
	Id           string  `json:"id"`
	Start_date   string  `json:"start_date"`
	Next_payment string  `json:"next_payment"`
	Status       string  `json:"status"`
	Plan         Plan    `json:"plan"`
	Charges      Charges `json:"charges"`
}

type Plan struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Frequency   string `json:"frequency"`
	Qty_charges int8   `json:"qty_charges"`
}

type Charges struct {
	Completed []Completed `json:"completed"`
	Future    []Future    `json:"future"`
}

type Completed struct {
	Order_id          string `json:"order_id"`
	Amount            string `json:"amount"`
	Status            string `json:"status"`
	Installments      int8   `json:"installments"`
	Card_type         string `json:"card_type"`
	Card_last_digits  string `json:"card_last_digits"`
	Card_first_digits string `json:"card_first_digits"`
	Created_at        string `json:"created_at"`
}

type Future struct {
	Charge_date string `json:"charge_date"`
}
