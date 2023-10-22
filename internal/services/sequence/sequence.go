package sequence

import "go.mongodb.org/mongo-driver/bson/primitive"

type Sequence struct {
	Id          *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string              `json:"name" bson:"name"`
	DisplayName string              `json:"displayName" bson:"displayName"`
	Duration    int                 `json:"duration" bson:"duration"`
	// Visible              bool      `json:"visible" gorm:"column:visible"`
	// Votes                int       `json:"votes" gorm:"column:votes"`
	// VoteTime             time.Time `json:"voteTime" gorm:"column:voteTime"`
	// VotesTotal           int       `json:"votesTotal" gorm:"column:votesTotal"`
	// Index                int       `json:"index" gorm:"column:index"`
	// Order                int       `json:"order" gorm:"column:0rder"`
	ImageUrl string `json:"imageUrl" bson:"imageUrl"`
	Active   bool   `json:"active" bson:"active"`
	Psa      bool   `json:"psa" bson:"psa"`

	// OwnerVoted           bool      `json:"ownerVoted" gorm:"column:ownerVoted"`
	// SequenceVisibleCount int       `json:"sequenceVisibleCount" gorm:"column:sequenceVisibleCount"`
	Type string `json:"type" bson:"type"`
	// GroupName            string    `json:"groupName" gorm:"column:groupName"`
}
