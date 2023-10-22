package mode

import "go.mongodb.org/mongo-driver/bson/primitive"

type Jukebox struct {
	Id                     *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NextPlaylist           string              `json:"nextPlaylist" bson:"nextPlaylist"`
	FuturePlaylist         string              `json:"futurePlaylist" bson:"futurePlaylist"`
	FuturePlaylistSequence int                 `json:"futurePlaylistSequence" bson:"futurePlaylistSequence"`
	OwnerRequested         bool                `json:"ownerRequested" bson:"ownerRequested"`
	Sequence               string              `json:"sequence" bson:"sequence"`
}
