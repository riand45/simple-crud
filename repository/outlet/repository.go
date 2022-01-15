package outlet

import "gorm.io/gorm"


type Repository interface {
	FindAll(userID int) ([]Outlet, error)
	FindByID(userID int, ID int) (Outlet, error)
	Save(outlet Outlet) (Outlet, error)
	Update(Outlet Outlet) (Outlet, error)
	Destroy(userID int, ID int) (Outlet, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(userID int) ([]Outlet, error) {
	var outlets []Outlet

	err := r.db.Where("user_id", userID).Find(&outlets).Error
	if err != nil {
		return outlets, err
	}

	return outlets, nil
}

func (r *repository) FindByID(userID int, ID int) (Outlet, error) {
	var outlet Outlet
	err := r.db.Where("user_id", userID).Where("id = ?", ID).Find(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (r *repository) Save(outlet Outlet) (Outlet, error) {
	err := r.db.Create(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (r *repository) Update(outlet Outlet) (Outlet, error) {
	err := r.db.Save(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (r *repository) Destroy(userID int, ID int) (Outlet, error) {
	var outlet Outlet
	err := r.db.Where("user_id", userID).Where("id = ?", ID).Delete(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}
