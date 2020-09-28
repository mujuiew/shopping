package structtype

// Input ...
type Input struct {
	IAccountID         string `json:"iaccount_id"`
	IAccountFirsteName string `json:"iaccount_first_name"`
	IAccountLastName   string `json:"iaccount_last_name"`
	IAccountEmail      string `json:"iaccount_email"`
	IAccountPhone      string `json:"iaccount_phone"`
}

// InputProdut ...
type InputProdut struct {
	IProAccount string `json:"iproduct_account"`
	IProName    string `json:"iproduct_name"`
	IProPrice   int    `json:"iproduct_price"`
	IProAmount  int    `json:"iproduct_amount"`
	IProImg     string `json:"iproduct_img"`
}

// Output ...
type Output struct {
	AccountID string `json:"account_id"`
	Status    string `json:"status"`
}

// OutputStore ...
type OutputStore struct {
	ProductID     string `json:"product_id"`
	ProductName   string `json:"product_name"`
	ProductImg    string `json:"product_image"`
	ProductPrice  int    `json:"product_price"`
	ProductAmount int    `json:"product_amount"`
}

// OutputPro ...
type OutputPro struct {
	StatusInsert error `json:"status_insert"`
}

// ProID ...
type ProID struct {
	ProductID     string `json:"product_id"`
	ProductName   string `json:"product_name"`
	ProductImg    string `json:"product_image"`
	ProductPrice  int    `json:"product_price"`
	ProductAmount int    `json:"product_amount"`
}

// ReAcc ...
type ReAcc struct {
	AccountID string `json:"acc_id"`
}

// LastID ...
var LastID ReAcc

// PROID ...
var PROID ProID
