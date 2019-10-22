package model

import (
	"github.com/jinzhu/gorm"
)

type Bookings []*Booking

type Booking struct {
	gorm.Model
	Course      string
	Class       string
	Lab         string
	Teacher     string
	Description string
}

type BookingDtos []*BookingDto

type BookingDto struct {
	ID          uint   `json:"id"`
	Course      string `json:"course"`
	Class       string `json:"class"`
	Lab         string `json:"lab"`
	Teacher     string `json:"teacher"`
	Description string `json:"description"`
}

func (b Booking) ToDto() *BookingDto {
	return &BookingDto{
		ID:          b.ID,
		Course:      b.Course,
		Class:       b.Class,
		Lab:         b.Lab,
		Teacher:     b.Teacher,
		Description: b.Description,
	}
}

func (bs Bookings) ToDto() BookingDtos {
	dtos := make([]*BookingDto, len(bs))
	for i, b := range bs {
		dtos[i] = b.ToDto()
	}
	return dtos
}

type BookingForm struct {
	Course      string `json:"course"`
	Class       string `json:"class"`
	Lab         string `json:"lab"`
	Teacher     string `json:"teacher"`
	Description string `json:"description"`
}

func (f *BookingForm) ToModel() (*Booking, error) {
	return &Booking{
		Course:      f.Course,
		Class:       f.Class,
		Lab:         f.Lab,
		Teacher:     f.Teacher,
		Description: f.Description,
	}, nil
}
