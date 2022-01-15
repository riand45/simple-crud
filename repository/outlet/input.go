package outlet

import ("majoo/repository/user")

type OutletAll struct {
	Name 			string `json:"name" binding:"required"`
	Description		string `json:"description" binding:"required"`
	User			user.User
}

type GetOutletDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateOutletInput struct {
	Name 			string `json:"name" binding:"required"`
	Description		string `json:"description" binding:"required"`
	User			user.User
}
