package remote

import (
	"github.com/rparmer/remote-falcon-api/internal/database"
)

type RemoteService struct {
	db *database.MongoDB
}

func NewService(db *database.MongoDB) *RemoteService {
	return &RemoteService{db: db}
}
