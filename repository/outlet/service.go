package outlet

import (
	"errors"
	"fmt"
)

type Service interface {
	FindOutlets(userID int) ([] Outlet, error)
	FindOutlet(userID int, input GetOutletDetailInput) (Outlet, error)
	CreateOutlet(input CreateOutletInput) (Outlet, error)
	UpdateOutlet(userID int, inputID GetOutletDetailInput, inputData CreateOutletInput) (Outlet, error)
	DeleteOutlet(userID int, input GetOutletDetailInput) (Outlet, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindOutlets(userID int) ([]Outlet, error) {
	outlets, err := s.repository.FindAll(userID)
	if err != nil {
		return outlets, err
	}

	return outlets, nil
}

func (s *service) FindOutlet(userID int, input GetOutletDetailInput) (Outlet, error) {
	outlet, err := s.repository.FindByID(userID, input.ID)

	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (s *service) CreateOutlet(input CreateOutletInput) (Outlet, error) {
	outlet := Outlet{}
	outlet.Name = input.Name
	outlet.Description = input.Description
	outlet.UserID = input.User.ID

	newOutlet, err := s.repository.Save(outlet)
	if err != nil {
		return newOutlet, err
	}

	return newOutlet, nil
}

func (s *service) UpdateOutlet(userID int, inputID GetOutletDetailInput, inputData CreateOutletInput) (Outlet, error) {
	outlet, err := s.repository.FindByID(userID, inputID.ID)
	if err != nil {
		return outlet, err
	}
	fmt.Println(userID)
	fmt.Println(inputID.ID)
	fmt.Println(outlet)

	if userID != inputData.User.ID {
		return outlet, errors.New("Not an owner of the outlet")
	}

	outlet.Name = inputData.Name
	outlet.Description = inputData.Description

	updatedOutlet, err := s.repository.Update(outlet)
	if err != nil {
		return updatedOutlet, err
	}

	return updatedOutlet, nil
}

func (s *service) DeleteOutlet(userID int, input GetOutletDetailInput) (Outlet, error) {
	outlet, err := s.repository.Destroy(userID, input.ID)

	if err != nil {
		return outlet, err
	}

	return outlet, nil
}