package viewer

import "time"

type ViewerJukeStats struct {
	ViewerJukeStatKey int64                              `json:"viewerJukeStatKey" gorm:"column:viewerJukeStatKey;autoIncrement;primaryKey"`
	RemoteToken       string                             `json:"remoteToken" gorm:"column:remoteToken"`
	PlaylistName      string                             `json:"playlistName" gorm:"column:playlistName"`
	RequestDateTime   time.Time                          `json:"requestDateTime" gorm:"column:requestDateTime"`
	RequestDate       int64                              `json:"requestDate"`
	TotalRequests     int                                `json:"totalRequests"`
	SequenceRequests  *[]ViewerJukeStatsSequenceRequests `json:"sequenceRequests"`
}

type ViewerJukeStatsSequenceRequests struct {
	SequenceName     string `json:"sequenceName"`
	SequenceRequests int    `json:"sequenceRequests"`
}

func (ViewerJukeStats) TableName() string {
	return "VIEWER_JUKE_STATS"
}

func (s *ViewerService) GetViewerJukeStats(id int64) (*ViewerJukeStats, error) {
	var v ViewerJukeStats
	result := s.db.First(&v, "remoteId = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &v, nil
}

func (s *ViewerService) SaveViewerJukeStats(r ViewerJukeStats) (int64, error) {
	result := s.db.Save(r)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
