package remote

import "time"

type PsaSequence struct {
	Id                    int64     `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
	RemoteId              string    `json:"remoteId" gorm:"column:remoteId"`
	PsaSequenceName       string    `json:"psaSequenceName" gorm:"column:psaSequenceName"`
	PsaSequenceOrder      int       `json:"psaSequenceOrder" gorm:"column:psaSequenceOrder"`
	PsaSequenceLastPlayed time.Time `json:"psaSequenceLastPlayed" gorm:"column:psaSequenceLastPlayed"`
}

func (PsaSequence) TableName() string {
	return "PSA_SEQUENCE"
}

// func (s *RemoteService) GetPsaSequence(id int64) (*PsaSequence, error) {
// 	var p PsaSequence
// 	result := s.db.First(&p, "remoteId = ?", id)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &p, nil
// }

// func (s *RemoteService) SavePsaSequence(p PsaSequence) (int64, error) {
// 	result := s.db.Save(p)
// 	if result.Error != nil {
// 		return 0, result.Error
// 	}
// 	return result.RowsAffected, nil
// }
