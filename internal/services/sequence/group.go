package sequence

import "go.mongodb.org/mongo-driver/bson/primitive"

type SequenceGroup struct {
	Id   *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string              `json:"name" bson:"name"`
	// Votes                     int       `json:"votes" gorm:"column:votes"`
	// VoteTime                  time.Time `json:"voteTime" gorm:"column:voteTime"`
	// VoteTotal                 int       `json:"voteTotal" gorm:"column:voteTotal"`
	// PlaylistsInGroup          int       `json:"playlistsInGroup" gorm:"column:playlistsInGroup"`
	// SequenceGroupVisibleCount int       `json:"sequenceGroupVisibleCount" gorm:"column:sequenceGroupVisibleCount"`
	Sequences *[]string `json:"sequences" bson:"sequences"`
}
