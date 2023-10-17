package viewer

import "time"

type ViewerVoteStats struct {
	ViewerVoteStatKey int64     `json:"viewerVoteStatKey" gorm:"column:viewerVoteStatKey;autoIncrement;primaryKey"`
	RemoteToken       string    `json:"remoteToken" gorm:"column:remoteToken"`
	PlaylistName      string    `json:"playlistName" gorm:"column:playlistName"`
	VoteDateTime      time.Time `json:"voteDateTime" gorm:"column:voteDateTime"`
}

type ViewerVoteStatsSequenceVotes struct {
	SequenceName  string `json:"sequenceName"`
	SequenceVotes int    `json:"sequenceVotes"`
}

func (ViewerVoteStats) TableName() string {
	return "VIEWER_VOTE_STATS"
}

func (s *ViewerService) GetViewerVoteStats(id int64) (*ViewerVoteStats, error) {
	var v ViewerVoteStats
	result := s.db.First(&v, "remoteId = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &v, nil
}

func (s *ViewerService) SaveViewerVoteStats(r ViewerPageStats) (int64, error) {
	result := s.db.Save(r)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
