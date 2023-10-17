package viewer

import "time"

type ViewerPageStats struct {
	ViewerPageStatKey int64     `json:"viewerPageStatKey" gorm:"column:viewerPageStatKey;autoIncrement;primaryKey"`
	RemoteToken       string    `json:"remoteToken" gorm:"column:remoteToken"`
	PageVisitIp       string    `json:"pageVisitIp" gorm:"column:pageVisitIp"`
	PageVisitDateTime time.Time `json:"pageVisitDateTime" gorm:"column:pageVisitDateTime"`
	PageVisitDate     int64     `json:"pageVisitDate"`
	UniqueVisits      int       `json:"uniqueVisits"`
	TotalVisits       int       `json:"totalVisits"`
}

func (ViewerPageStats) TableName() string {
	return "VIEWER_PAGE_STATS"
}

func (s *ViewerService) GetViewerPageStats(id int64) (*ViewerPageStats, error) {
	var v ViewerPageStats
	result := s.db.First(&v, "remoteId = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &v, nil
}

func (s *ViewerService) SaveViewerPageStats(r ViewerPageStats) (int64, error) {
	result := s.db.Save(r)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
