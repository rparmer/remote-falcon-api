package playlist

import (
	"time"
)

type Playlist struct {
	Id                   int64     `json:"id" gorm:"column:Id;autoIncrement;primaryKey"`
	RemoteId             string    `json:"remoteId" gorm:"column:remoteId"`
	Name                 string    `json:"name" gorm:"column:name"`
	PrettyName           string    `json:"prettyName" gorm:"column:prettyName"`
	Duration             int       `json:"duration" gorm:"column:duration"`
	Visible              bool      `json:"visible" gorm:"column:visible"`
	Votes                int       `json:"votes" gorm:"column:votes"`
	VoteTime             time.Time `json:"voteTime" gorm:"column:voteTime"`
	VotesTotal           int       `json:"votesTotal" gorm:"column:votesTotal"`
	Index                int       `json:"index" gorm:"column:index"`
	Order                int       `json:"order" gorm:"column:0rder"`
	ImageUrl             string    `json:"imageUrl" gorm:"column:imageUrl"`
	IsActive             bool      `json:"isActive" gorm:"column:isActive"`
	OwnerVoted           bool      `json:"ownerVoted" gorm:"column:ownerVoted"`
	SequenceVisibleCount int       `json:"sequenceVisibleCount" gorm:"column:sequenceVisibleCount"`
	Type                 string    `json:"type" gorm:"column:type"`
	GroupName            string    `json:"groupName" gorm:"column:groupName"`
}

type SyncPlaylistDetails struct {
	Name     string `json:"name"`
	Duration int    `json:"duration"`
	Index    int    `json:"index"`
	Type     string `json:"type"`
}

func (Playlist) TableName() string {
	return "PLAYLIST"
}

func (s *PlaylistService) GetPlaylist(id int) (*Playlist, error) {
	var p Playlist
	result := s.db.First(&p, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &p, nil
}

func (s *PlaylistService) GetAllPlaylists(id int) (*[]Playlist, error) {
	var p []Playlist
	result := s.db.Find(&p, "remoteId = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &p, nil
}

func (s *PlaylistService) SavePlaylist(p Playlist) (int64, error) {
	result := s.db.Save(p)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (s *PlaylistService) DeletePlaylist(p Playlist) (int64, error) {
	result := s.db.Delete(p)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
