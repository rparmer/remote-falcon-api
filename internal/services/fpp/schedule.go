package fpp

type FppSchedule struct {
	Id                    int64  `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
	NextScheduledSequence string `json:"nextScheduledSequence" gorm:"column:nextScheduledSequence"`
}

func (FppSchedule) TableName() string {
	return "FPP_SCHEDULE"
}

func (s *FppService) GetFppSchedule(id int) (*FppSchedule, error) {
	var p FppSchedule
	result := s.db.First(&p, "remoteId = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &p, nil
}

func (s *FppService) SaveFppSchedule(p FppSchedule) (int64, error) {
	result := s.db.Save(p)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
