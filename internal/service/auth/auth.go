package auth

import (
	"fmt"

	"github.com/rparmer/remote-falcon-api/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	authCollection = "auths"
)

type Service struct {
	db *database.MongoDB
}

type Auth struct {
	Id          *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	PluginToken string              `json:"pluginToken" bson:"pluginToken"`
	Username    string              `json:"username" bson:"username"`
	Password    string              `json:"password" bson:"password"`
}

func New(db *database.MongoDB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateAuth(a Auth) error {
	if err := s.db.Create(authCollection, &a); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAuth() (*Auth, error) {
	var auth Auth
	if err := s.db.Get(authCollection, bson.D{}, &auth); err != nil {
		return nil, err
	}
	return &auth, nil
}

func (s *Service) UpdateToken(id primitive.ObjectID, token string) error {
	count, err := s.db.Update(authCollection, bson.D{{Key: "_id", Value: id}}, bson.D{{Key: "token", Value: token}})
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no records found")
	}
	return nil
}
