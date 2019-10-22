package repository

import (
	"bass/model"

	"github.com/jinzhu/gorm"
)

func ListBookings(db *gorm.DB) (model.Bookings, error) {
	bookings := make([]*model.Booking, 0)
	if err := db.Find(&bookings).Error; err != nil {
		return nil, err
	}

	return bookings, nil
}

func ReadBooking(db *gorm.DB, id uint) (*model.Booking, error) {
	booking := &model.Booking{}
	if err := db.Where("id = ?", id).First(&booking).Error; err != nil {
		return nil, err
	}

	return booking, nil
}

func DeleteBooking(db *gorm.DB, id uint) error {
	booking := &model.Booking{}
	if err := db.Where("id = ?", id).Delete(&booking).Error; err != nil {
		return err
	}

	return nil
}

func CreateBooking(db *gorm.DB, booking *model.Booking) (*model.Booking, error) {
	if err := db.Create(booking).Error; err != nil {
		return nil, err
	}

	return booking, nil
}

func UpdateBooking(db *gorm.DB, booking *model.Booking) error {
	if err := db.First(&model.Booking{}, booking.ID).Update(booking).Error; err != nil {
		return err
	}

	return nil
}
