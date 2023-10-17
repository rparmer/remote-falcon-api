package fpp

type FppStats struct {
	Id                     int64  `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
	RemoteId               int64  `json:"remoteId" gorm:"column:remoteId"`
	FppdStatus             string `json:"fppdStatus" gorm:"column:fppdStatus"`
	FppStatus              string `json:"fppStatus" gorm:"column:fppStatus"`
	CpuTemp                string `json:"cpuTemp" gorm:"column:cpuTemp"`
	Volume                 int    `json:"volume" gorm:"column:volume"`
	CurrentPlayingSequence string `json:"currentPlayingSequence" gorm:"column:currentPlayingSequence"`
	CurrentPlayingPlaylist string `json:"currentPlayingPlaylist" gorm:"column:currentPlayingPlaylist"`
	ScheduledPlaylist      string `json:"scheduledPlaylist" gorm:"column:scheduledPlaylist"`
	ScheduledStartTime     string `json:"scheduledStartTime" gorm:"column:scheduledStartTime"`
	ScheduledEndTime       string `json:"scheduledEndTime" gorm:"column:scheduledEndTime"`
}

func (FppStats) TableName() string {
	return "FPP_STATS"
}

func (s *FppService) GetFppStats(id int) (*FppStats, error) {
	var p FppStats
	result := s.db.First(&p, "remoteId = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &p, nil
}

func (s *FppService) SaveFppStats(p FppStats) (int64, error) {
	result := s.db.Save(p)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
