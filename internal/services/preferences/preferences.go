package preferences

import "go.mongodb.org/mongo-driver/bson/primitive"

type Preferences struct {
	Id                   *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Token                string              `json:"token" bson:"token"`
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

	// ViewerPagePublic             bool                       `json:"viewerPagePublic" gorm:"column:viewerPagePublic"`
	// PsaSequence                  string                     `json:"psaSequence" gorm:"column:psaSequence"`
	// AutoSwitchControlModeSize    int                        `json:"autoSwitchControlModeSize" gorm:"column:autoSwitchControlModeSize"`
	// AutoSwitchControlModeToggled bool                       `json:"autoSwitchControlModeToggled" gorm:"column:autoSwitchControlModeToggled"`
	// HideSequenceCount            int                        `json:"hideSequenceCount" gorm:"column:hideSequenceCount"`

}
