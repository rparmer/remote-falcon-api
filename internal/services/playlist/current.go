package playlist

type CurrentPlaylist struct {
	Id                int64  `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
	RemoteId          int64  `json:"remoteId" gorm:"column:remoteId"`
	CurrentPlaylistId string `json:"currentPlaylistId" gorm:"column:currentPlaylistId"`
}

func (CurrentPlaylist) TableName() string {
	return "CURRENT_PLAYLIST"
}

func (s *PlaylistService) GetCurrentPlaylist(id int) (*CurrentPlaylist, error) {
	var p CurrentPlaylist
	result := s.db.First(&p, "remoteId = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &p, nil
}

func (s *PlaylistService) SaveCurrentPlaylist(p CurrentPlaylist) (int64, error) {
	result := s.db.Save(p)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
