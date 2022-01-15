package outlet

import ("time")

type Outlet struct {
	ID 				int
	UserID 			int
	Name 			string
	Description		string
	CreatedAt		time.Time
	UpdatedAt		time.Time
}
