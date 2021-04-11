package rest

import "../storage"

// accountRequest represents request for creating account
type userRequest struct {
	UserID string `json:"user_id"`
	Mobile string `json:"mobile_no"`
}
type carRequest struct {
	ID               string
	carLicenseNumber string
	manufacturer     string
	model            string
	basePrice        int
	PPH              int
	securitydeposit  int
	available        bool
}

// mapToModel maps request to dao model
func (request userRequest) mapToModel() storage.User {
	var user storage.User
	user.ID = request.UserID
	user.Mobile = request.Mobile
	return user
}
func (request carRequest) DAO() storage.Car {
	var car storage.Car
	car.ID = request.ID
	car.CarLicenseNumber = request.carLicenseNumber
	car.Manufacturer = request.manufacturer
	car.Model = request.model
	car.BasePrice = request.basePrice
	car.PPH = request.PPH
	car.Securitydeposit = request.securitydeposit
	car.Available = request.available
	return car

}
func (request carRequest) DAOcarBooking() storage.carBooking {
	var car storage.carBooking
	car.BookingId = request.ID
	car.CarLicenseNumber = request.carLicenseNumber
	car.Manufacturer = request.manufacturer
	car.Model = request.model
	car.BasePrice = request.basePrice
	car.PPH = request.PPH
	car.Securitydeposit = request.securitydeposit
	car.Available = request.available
	return car

}
