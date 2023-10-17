package playlist

import "gorm.io/gorm"

type PlaylistService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *PlaylistService {
	return &PlaylistService{db: db}
}
