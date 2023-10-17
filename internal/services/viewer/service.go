package viewer

import "gorm.io/gorm"

type ViewerService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *ViewerService {
	return &ViewerService{db: db}
}
