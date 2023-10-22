package service

import "github.com/rparmer/remote-falcon-api/internal/database"

type Service struct {
	db *database.MongoDB
}

func New(db *database.MongoDB) *Service {
	return &Service{db: db}
}
