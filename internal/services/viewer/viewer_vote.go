package viewer

type RemoteViewerVote struct {
	Id       int64  `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
	RemoteId string `json:"remoteId" gorm:"column:remoteId"`
	ViewerIp string `json:"viewerIp" gorm:"column:viewerIp"`
}

func (RemoteViewerVote) TableName() string {
	return "REMOTE_VIEWER_VOTES"
}

func (s *ViewerService) GetViewerVote(id int64) (*RemoteViewerVote, error) {
	var r RemoteViewerVote
	result := s.db.First(&r, "remoteId = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &r, nil
}

func (s *ViewerService) SaveViewerVote(r RemoteViewerVote) (int64, error) {
	result := s.db.Save(r)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
