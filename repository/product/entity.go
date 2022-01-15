package product

import ("time")

type Product struct {
	ID 				int
	OutletID 		int
	Name 			string
	Description		string
	Stock			int
	Price			int
	CreatedAt		time.Time
	UpdatedAt		time.Time
}

type ProductImage struct {
	ID         			int
	ProductID 			int
	ProductFileName   	string
	CreatedAt  			time.Time
	UpdatedAt  			time.Time
}
