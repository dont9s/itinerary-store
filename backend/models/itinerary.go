package models

import "gorm.io/gorm"

type Itinerary struct {
    gorm.Model
    Title string
    Description string
    ImageURL string
    Duration string
    Price int
    FilePath string
}
