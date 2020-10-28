package models

type Orders struct {
	PaymentType        string             `json:"payment_type"`
	TransactionDetails TransactionDetails `json:"transaction_details"`
	ItemDetails        []ItemDetails      `json:"item_details"`
	CustomerDetails    CustomerDetails    `json:"customer_details"`
	Gopay              Gopay              `json:"gopay"`
}

type TransactionDetails struct {
	OrderID     string `json:"order_id"`
	GrossAmount int    `json:"gross_amount"`
}

type ItemDetails struct {
	ID       string `json:"id"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Name     string `json:"name"`
}

type CustomerDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type Gopay struct {
	EnableCallback bool   `json:"enable_callback"`
	CallbackURL    string `json:"callback_url"`
}

// {
//   "payment_type": "gopay",
//   "transaction_details": {
//     "order_id": "order03",
//     "gross_amount": 275000
//   },
//   "item_details": [
//     {
//       "id": "id1",
//       "price": 275000,
//       "quantity": 1,
//       "name": "Bluedio H+ Turbine Headphone with Bluetooth 4.1 -"
//     }
//   ],
//   "customer_details": {
//     "first_name": "Budi",
//     "last_name": "Utomo",
//     "email": "budi.utomo@midtrans.com",
//     "phone": "081223323423"
//   },
//   "gopay": {
//     "enable_callback": true,
//     "callback_url": "someapps://callback"
//   }
// }
