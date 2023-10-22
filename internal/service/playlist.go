package service

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Playlist struct {
	Id                   *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name                 string              `json:"name" bson:"name"`
	PrettyName           string              `json:"prettyName" bson:"prettyName"`
	Duration             int                 `json:"duration" bson:"duration"`
	Visible              bool                `json:"visible" bson:"visible"`
	Votes                int                 `json:"votes" bson:"votes"`
	VoteTime             time.Time           `json:"voteTime" bson:"voteTime"`
	VotesTotal           int                 `json:"votesTotal" bson:"votesTotal"`
	Index                int                 `json:"index" bson:"index"`
	Order                int                 `json:"order" bson:"0rder"`
	ImageUrl             string              `json:"imageUrl" bson:"imageUrl"`
	IsActive             bool                `json:"isActive" bson:"isActive"`
	OwnerVoted           bool                `json:"ownerVoted" bson:"ownerVoted"`
	SequenceVisibleCount int                 `json:"sequenceVisibleCount" bson:"sequenceVisibleCount"`
	Type                 string              `json:"type" bson:"type"`
	GroupName            string              `json:"groupName" bson:"groupName"`
	Current              bool                `json:"current" bson:"current"`
}

// type CurrentPlaylist struct {
// 	Id                *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
// 	RemoteId          int64               `json:"remoteId" bson:"remoteId"`
// 	CurrentPlaylistId string              `json:"currentPlaylistId" bson:"currentPlaylistId"`
// }

type PlaylistGroup struct {
	Id                        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	RemoteId                  string              `json:"remoteId" bson:"remoteId"`
	Name                      string              `json:"name" bson:"name"`
	Votes                     int                 `json:"votes" bson:"votes"`
	VoteTime                  time.Time           `json:"voteTime" bson:"voteTime"`
	VoteTotal                 int                 `json:"voteTotal" bson:"voteTotal"`
	PlaylistsInGroup          int                 `json:"playlistsInGroup" bson:"playlistsInGroup"`
	SequenceGroupVisibleCount int                 `json:"sequenceGroupVisibleCount" bson:"sequenceGroupVisibleCount"`
	SequenceNamesInGroup      *[]string           `json:"sequenceNamesInGroup" bson:"sequenceNamesInGroup"`
}
