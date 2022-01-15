package product

import (
	"majoo/repository/user"
	"majoo/repository/outlet"
)

type ProductAll struct {
	Name 			string `json:"name" binding:"required"`
	OutletID		int `json:"outlet_id" binding:"required"`
	Description		string `json:"description" binding:"required"`
	Stock			int `json:"stock" binding:"required"`
	Price			int `json:"price" binding:"required"`
	Outlet			outlet.Outlet
}

type GetOutletProductInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type GetOutletDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateProductInput struct {
	Name 			string `json:"name" binding:"required"`
	OutletID		int `json:"outlet_id" binding:"required"`
	Description		string `json:"description" binding:"required"`
	Stock			int `json:"stock" binding:"required"`
	Price			int `json:"price" binding:"required"`
}

type CreateProductImageInput struct {
	ProductID 	int  `form:"product_id" binding:"required"`
	// ProductFileName string `form:"product_file_name" binding:"required"`
}
