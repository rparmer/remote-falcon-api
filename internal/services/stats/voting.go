package stats

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vote struct {
	Id            *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	PlaylistName  string              `json:"playlistName" borm:"playlistName"`
	VoteDate      time.Time           `json:"voteDateTime" borm:"voteDate"`
	SequenceName  string              `json:"sequenceName"`
	SequenceVotes int                 `json:"sequenceVotes"`
	ViewerIp      string              `json:"viewerIp" bson:"viewerIp"`
}
