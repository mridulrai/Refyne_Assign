package storage

import (
	"time"
)

// User represents User table fields
type User struct {
	ID     string
	Mobile string
	Active bool
}

// CAR represents car table fields
type Car struct {
	ID               string
	CarLicenseNumber string
	Manufacturer     string
	Model            string
	BasePrice        int
	PPH              int
	Securitydeposit  int
	Available        bool
}

// carBooking represents carBooking table fields
type carBooking struct {
	BookingId     string
	CarID         string
	UserID        string
	StartDateTime *time.Time
	EndDateTime   *time.Time
}
