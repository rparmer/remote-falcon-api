package playlist

import "time"

type PlaylistGroup struct {
	Id                        int64     `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
	RemoteId                  string    `json:"remoteId" gorm:"column:remoteId"`
	Name                      string    `json:"name" gorm:"column:name"`
	Votes                     int       `json:"votes" gorm:"column:votes"`
	VoteTime                  time.Time `json:"voteTime" gorm:"column:voteTime"`
	VoteTotal                 int       `json:"voteTotal" gorm:"column:voteTotal"`
	PlaylistsInGroup          int       `json:"playlistsInGroup" gorm:"column:playlistsInGroup"`
	SequenceGroupVisibleCount int       `json:"sequenceGroupVisibleCount" gorm:"column:sequenceGroupVisibleCount"`
	SequenceNamesInGroup      *[]string `json:"sequenceNamesInGroup"`
}

func (PlaylistGroup) TableName() string {
	return "PLAYLIST_GROUP"
}

func (s *PlaylistService) GetPlaylistGroup(id int) (*PlaylistGroup, error) {
	var p PlaylistGroup
	result := s.db.First(&p, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &p, nil
}

func (s *PlaylistService) SavePlaylistGroup(p PlaylistGroup) (int64, error) {
	result := s.db.Save(p)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (s *PlaylistService) DeletePlaylistGroup(p PlaylistGroup) (int64, error) {
	result := s.db.Delete(p)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
