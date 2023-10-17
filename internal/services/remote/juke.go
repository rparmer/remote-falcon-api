package remote

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const jukeCollection = "remote_jukes"

type RemoteJuke struct {
	Id                     *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RemoteId               int                 `json:"remoteId" bson:"remoteId"`
	NextPlaylist           string              `json:"nextPlaylist" bson:"nextPlaylist"`
	FuturePlaylist         string              `json:"futurePlaylist" bson:"futurePlaylist"`
	FuturePlaylistSequence int                 `json:"futurePlaylistSequence" bson:"futurePlaylistSequence"`
	OwnerRequested         bool                `json:"ownerRequested" bson:"ownerRequested"`
	Sequence               string              `json:"sequence" bson:"sequence"`
}

func (s *RemoteService) GetJuke(id int) (*RemoteJuke, error) {
	var r RemoteJuke
	err := s.db.Get(jukeCollection, bson.D{{Key: "remoteId", Value: id}}, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *RemoteService) SaveJuke(r RemoteJuke) error {
	err := s.db.Save(jukeCollection, r)
	if err != nil {
		return err
	}
	return nil
}
