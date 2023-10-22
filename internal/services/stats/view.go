package stats

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type View struct {
	Id              *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ViewerIp        string              `json:"viewerIp" bson:"viewerIp"`
	FirstVisitDate  time.Time           `json:"firstVisitDate" bson:"firstVisitDate"`
	LatestVisitDate time.Time           `json:"latestVisitDate" bson:"latestVisitDate"`
}
