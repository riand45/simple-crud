package product

type ProductFormatter struct {
	ID 				int `json:"id"`
	Name 			string `json:"name"`
	OutletID		int `json:"outlet_id"`
	Description 	string `json:"description"`
	Stock 			int `json:"stock"`
	Price 			int `json:"price"`
}