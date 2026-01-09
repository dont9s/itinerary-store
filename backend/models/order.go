package models

import "gorm.io/gorm"

type Order struct {
    gorm.Model
    ItineraryID uint
    Amount int
    RazorpayOrder string
    Paid bool
    DownloadToken string
}
