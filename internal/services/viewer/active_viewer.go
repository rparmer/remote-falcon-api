package viewer

import (
	"time"
)

type ActiveViewer struct {
	Id                 int64     `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
	RemoteId           int64     `json:"remoteId" gorm:"column:remoteId"`
	ViewerIp           string    `json:"viewerIp" gorm:"column:viewerIp"`
	LastUpdateDateTime time.Time `json:"lastUpdateDateTime" gorm:"column:lastUpdateDateTime"`
}

func (ActiveViewer) TableName() string {
	return "ACTIVE_VIEWER"
}

// type ActiveViewerService struct {
// 	db *gorm.DB
// }

// func NewService(db *gorm.DB) *ActiveViewerService {
// 	return &ActiveViewerService{db: db}
// }

func (s *ViewerService) GetActiveViewer(id int64) (*ActiveViewer, error) {
	var v ActiveViewer
	result := s.db.First(&v, "remoteId = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &v, nil
}

func (s *ViewerService) SaveActiveViewer(v ActiveViewer) (int64, error) {
	result := s.db.Save(v)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// func (s *ViewerService) DeleteAllByRemoteToken(token string) error {
// 	result := s.db.Delete(&ActiveViewer{}, "remoteToken = ?", token)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }

// func (s *ViewerService) FindAllByRemoteToken(token string) (*[]ActiveViewer, error) {
// 	var a []ActiveViewer
// 	result := s.db.Find(&a, "remoteToken = ?", token)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &a, nil
// }

// func (s *ViewerService) FindFirstByRemoteTokenAndViewerIp(token string, ip string) (*ActiveViewer, error) {
// 	var a ActiveViewer
// 	result := s.db.First(&a, "remoteToken = ? AND viewerIp = ?", token, ip)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &a, nil
// }
