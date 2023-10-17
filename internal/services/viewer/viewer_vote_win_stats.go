package viewer

import "time"

type ViewerVoteWinStats struct {
	ViewerVoteWinStatKey int64                             `json:"viewerVoteWinStatKey" gorm:"column:viewerVoteWinStatKey;autoIncrement;primaryKey"`
	RemoteToken          string                            `json:"remoteToken" gorm:"column:remoteToken"`
	PlaylistName         string                            `json:"playlistName" gorm:"column:playlistName"`
	VoteWinDateTime      time.Time                         `json:"voteWinDateTime" gorm:"column:voteWinDateTime"`
	TotalVotes           int                               `json:"totalVotes" gorm:"column:totalVotes"`
	VoteDate             int64                             `json:"voteDate"`
	SequenceWins         *[]ViewerVoteWinStatsSequenceWins `json:"sequenceWins"`
}

type ViewerVoteWinStatsSequenceWins struct {
	SequenceName string `json:"sequenceName"`
	SequenceWins int    `json:"sequenceWins"`
}

func (ViewerVoteWinStats) TableName() string {
	return "VIEWER_VOTE_WIN_STATS"
}

func (s *ViewerService) GetViewerVoteWinStats(id int64) (*ViewerVoteWinStats, error) {
	var v ViewerVoteWinStats
	result := s.db.First(&v, "remoteId = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &v, nil
}

func (s *ViewerService) SaveViewerVoteWinStats(r ViewerVoteWinStats) (int64, error) {
	result := s.db.Save(r)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
