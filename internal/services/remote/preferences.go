package remote

import "github.com/rparmer/remote-falcon-api/internal/services/viewer"

type RemotePreferences struct {
	Id                           int64                      `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
	RemoteId                     int64                      `json:"remoteId" gorm:"column:remoteId"`
	ViewerModeEnabled            bool                       `json:"viewerModeEnabled" gorm:"column:viewerModeEnabled"`
	ViewerControlEnabled         bool                       `json:"viewerControlEnabled" gorm:"column:viewerControlEnabled"`
	ViewerControlMode            string                     `json:"viewerControlMode" gorm:"column:viewerControlMode"`
	ResetVotes                   bool                       `json:"resetVotes" gorm:"column:resetVotes"`
	JukeboxDepth                 int                        `json:"jukeboxDepth" gorm:"column:jukeboxDepth"`
	EnableGeolocation            bool                       `json:"enableGeolocation" gorm:"column:enableGeolocation"`
	RemoteLatitude               float32                    `json:"remoteLatitude" gorm:"column:remoteLatitude"`
	RemoteLongitude              float32                    `json:"remoteLongitude" gorm:"column:remoteLongitude"`
	AllowedRadius                float32                    `json:"allowedRadius" gorm:"column:allowedRadius"`
	MessageDisplayTime           int                        `json:"messageDisplayTime" gorm:"column:messageDisplayTime"`
	CheckIfVoted                 bool                       `json:"checkIfVoted" gorm:"column:checkIfVoted"`
	InterruptSchedule            bool                       `json:"interruptSchedule" gorm:"column:interruptSchedule"`
	PsaEnabled                   bool                       `json:"psaEnabled" gorm:"column:psaEnabled"`
	PsaSequence                  string                     `json:"psaSequence" gorm:"column:psaSequence"`
	PsaFrequency                 int                        `json:"psaFrequency" gorm:"column:psaFrequency"`
	JukeboxRequestLimit          int                        `json:"jukeboxRequestLimit" gorm:"column:jukeboxRequestLimit"`
	ViewerPagePublic             bool                       `json:"viewerPagePublic" gorm:"column:viewerPagePublic"`
	EnableLocationCode           bool                       `json:"enableLocationCode" gorm:"column:enableLocationCode"`
	LocationCode                 string                     `json:"locationCode" gorm:"column:locationCode"`
	AutoSwitchControlModeSize    int                        `json:"autoSwitchControlModeSize" gorm:"column:autoSwitchControlModeSize"`
	AutoSwitchControlModeToggled bool                       `json:"autoSwitchControlModeToggled" gorm:"column:autoSwitchControlModeToggled"`
	HideSequenceCount            int                        `json:"hideSequenceCount" gorm:"column:hideSequenceCount"`
	JukeboxHistoryLimit          int                        `json:"jukeboxHistoryLimit" gorm:"column:jukeboxHistoryLimit"`
	MakeItSnow                   bool                       `json:"makeItSnow" gorm:"column:makeItSnow"`
	ActiveTheme                  string                     `json:"activeTheme" gorm:"column:activeTheme"`
	HtmlContent                  string                     `json:"htmlContent" gorm:"column:htmlContent;type:mediumtext"`
	PsaSequenceList              *[]PsaSequence             `json:"psaSequenceList" gorm:"-"`
	RemoteViewerPages            *[]viewer.RemoteViewerPage `json:"remoteViewerPages" gorm:"-"`
	// ActiveRemoteViewerPage       string         `json:"activeRemoteViewerPage"`

	// ManagePsa                    string             `json:"managePsa" gorm:"column:managePsa"`
}

func (RemotePreferences) TableName() string {
	return "REMOTE_PREFERENCES"
}

// func (s *RemoteService) GetPreferences(id int64) (*RemotePreferences, error) {
// 	var r RemotePreferences
// 	result := s.db.First(&r, "remoteId = ?", id)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &r, nil
// }

// func (s *RemoteService) SavePreferences(r RemotePreferences) (int64, error) {
// 	result := s.db.Save(r)
// 	if result.Error != nil {
// 		return 0, result.Error
// 	}
// 	return result.RowsAffected, nil
// }
