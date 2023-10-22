package service

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	preferencesCollection = "preferences"
)

type Preferences struct {
	Id                   *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ViewerModeEnabled    bool                `json:"viewerModeEnabled" bson:"viewerModeEnabled"`
	ViewerControlEnabled bool                `json:"viewerControlEnabled" bson:"viewerControlEnabled"`
	ViewerControlMode    string              `json:"viewerControlMode" bson:"viewerControlMode"`
	EnableLocationCode   bool                `json:"enableLocationCode" bson:"enableLocationCode"`
	LocationCode         string              `json:"locationCode" bson:"locationCode"`
	EnableGeolocation    bool                `json:"enableGeolocation" bson:"enableGeolocation"`
	RemoteLatitude       float32             `json:"remoteLatitude" bson:"remoteLatitude"`
	RemoteLongitude      float32             `json:"remoteLongitude" bson:"remoteLongitude"`
	AllowedRadius        float32             `json:"allowedRadius" bson:"allowedRadius"`
	MessageDisplayTime   int                 `json:"messageDisplayTime" bson:"messageDisplayTime"`
	CheckIfVoted         bool                `json:"checkIfVoted" bson:"checkIfVoted"`
	InterruptSchedule    bool                `json:"interruptSchedule" bson:"interruptSchedule"`
	PsaEnabled           bool                `json:"psaEnabled" bson:"psaEnabled"`
	PsaFrequency         int                 `json:"psaFrequency" bson:"psaFrequency"`
	JukeboxDepth         int                 `json:"jukeboxDepth" bson:"jukeboxDepth"`
	JukeboxRequestLimit  int                 `json:"jukeboxRequestLimit" bson:"jukeboxRequestLimit"`
	JukeboxHistoryLimit  int                 `json:"jukeboxHistoryLimit" bson:"jukeboxHistoryLimit"`
	MakeItSnow           bool                `json:"makeItSnow" bson:"makeItSnow"`
	ActiveTheme          string              `json:"activeTheme" bson:"activeTheme"`
}

func (s *Service) CreatePreferences(p Preferences) error {
	if err := s.db.Create(preferencesCollection, &p); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdatePreferences(p Preferences) error {
	count, err := s.db.Update(preferencesCollection, bson.D{{Key: "_id", Value: p.Id}}, &p)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no records found")
	}
	return nil
}

func (s *Service) GetPreferences() (*Preferences, error) {
	var preferences Preferences
	if err := s.db.Get(preferencesCollection, bson.D{}, &preferences); err != nil {
		return nil, err
	}
	return &preferences, nil
}

func (s *Service) DeletePreferences(p Preferences) error {
	if err := s.db.Delete(preferencesCollection, bson.D{{Key: "_id", Value: p.Id}}); err != nil {
		return err
	}
	return nil
}
