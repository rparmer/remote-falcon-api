package sequence

import "github.com/rparmer/remote-falcon-api/internal/database"

type SequenceService struct {
	db *database.MongoDB
}

func NewService(db *database.MongoDB) *SequenceService {
	return &SequenceService{db: db}
}
