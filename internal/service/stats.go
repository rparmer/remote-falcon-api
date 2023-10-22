package service

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	jukeboxStatCollection = "jukeboxStats"
	votingStatCollection  = "votingStats"
	viewerStatCollection  = "viewerStats"
)

type JukeboxStat struct {
	Id            *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	PlaylistName  string              `json:"playlistName" borm:"playlistName"`
	VoteDate      time.Time           `json:"voteDateTime" borm:"voteDate"`
	SequenceName  string              `json:"sequenceName"`
	SequenceVotes int                 `json:"sequenceVotes"`
	ViewerIp      string              `json:"viewerIp" bson:"viewerIp"`
}

type VotingStat struct {
	Id            *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	PlaylistName  string              `json:"playlistName" borm:"playlistName"`
	VoteDate      time.Time           `json:"voteDateTime" borm:"voteDate"`
	SequenceName  string              `json:"sequenceName"`
	SequenceVotes int                 `json:"sequenceVotes"`
	ViewerIp      string              `json:"viewerIp" bson:"viewerIp"`
}

type ViewerStat struct {
	Id        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ViewerIp  string              `json:"viewerIp" bson:"viewerIp"`
	VisitDate time.Time           `json:"visitDate" bson:"visitDate"`
}

func (s *Service) CreateJukeboxStat(j JukeboxStat) error {
	if err := s.db.Create(jukeboxStatCollection, &j); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateJukeboxStat(j JukeboxStat) error {
	count, err := s.db.Update(jukeboxStatCollection, bson.D{{Key: "_id", Value: j.Id}}, &j)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no records found")
	}
	return nil
}

func (s *Service) GetJukeboxStat(name string) (*JukeboxStat, error) {
	var stat JukeboxStat
	if err := s.db.Get(jukeboxStatCollection, bson.D{{Key: "playlistName", Value: name}}, &stat); err != nil {
		return nil, err
	}
	return &stat, nil
}

func (s *Service) ListJukeboxStats() (*[]JukeboxStat, error) {
	var stats []JukeboxStat
	if err := s.db.List(jukeboxStatCollection, bson.D{}, &stats); err != nil {
		return nil, err
	}
	return &stats, nil
}

func (s *Service) DeleteJukeboxStat(j JukeboxStat) error {
	if err := s.db.Delete(jukeboxStatCollection, bson.D{{Key: "_id", Value: j.Id}}); err != nil {
		return err
	}
	return nil
}

func (s *Service) CreateVotingStat(v VotingStat) error {
	if err := s.db.Create(votingStatCollection, &v); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateVotingStat(v VotingStat) error {
	count, err := s.db.Update(votingStatCollection, bson.D{{Key: "_id", Value: v.Id}}, &v)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no records found")
	}
	return nil
}

func (s *Service) GetVotingStat(name string) (*VotingStat, error) {
	var stat VotingStat
	if err := s.db.Get(votingStatCollection, bson.D{{Key: "playlistName", Value: name}}, &stat); err != nil {
		return nil, err
	}
	return &stat, nil
}

func (s *Service) ListVotingStats() (*[]VotingStat, error) {
	var stats []VotingStat
	if err := s.db.List(votingStatCollection, bson.D{}, &stats); err != nil {
		return nil, err
	}
	return &stats, nil
}

func (s *Service) DeleteVotingStat(v VotingStat) error {
	if err := s.db.Delete(votingStatCollection, bson.D{{Key: "_id", Value: v.Id}}); err != nil {
		return err
	}
	return nil
}

func (s *Service) CreateViewerStat(v ViewerStat) error {
	if err := s.db.Create(viewerStatCollection, &v); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateViewerStat(v ViewerStat) error {
	count, err := s.db.Update(viewerStatCollection, bson.D{{Key: "_id", Value: v.Id}}, &v)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no records found")
	}
	return nil
}

func (s *Service) GetViewerStat(name string) (*ViewerStat, error) {
	var stat ViewerStat
	if err := s.db.Get(viewerStatCollection, bson.D{{Key: "playlistName", Value: name}}, &stat); err != nil {
		return nil, err
	}
	return &stat, nil
}

func (s *Service) ListViewerStats() (*[]ViewerStat, error) {
	var stats []ViewerStat
	if err := s.db.List(viewerStatCollection, bson.D{}, &stats); err != nil {
		return nil, err
	}
	return &stats, nil
}

func (s *Service) DeleteViewerStat(v ViewerStat) error {
	if err := s.db.Delete(viewerStatCollection, bson.D{{Key: "_id", Value: v.Id}}); err != nil {
		return err
	}
	return nil
}
