package viewer

import (
	"time"
)

type DefaultViewerPage struct {
	Id                int64     `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
	Version           float32   `json:"version" gorm:"column:version"`
	VersionCreateDate time.Time `json:"versionCreateDate" gorm:"column:versionCreateDate"`
	HtmlContent       string    `json:"htmlContent" gorm:"column:htmlContent"`
	IsVersionActive   bool      `json:"isVersionActive" gorm:"column:isVersionActive"`
}

func (DefaultViewerPage) TableName() string {
	return "DEFAULT_VIEWER_PAGE"
}

func (s *ViewerService) GetDefaultViewerPage(id int64) (*DefaultViewerPage, error) {
	var v DefaultViewerPage
	result := s.db.First(&v, "remoteId = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &v, nil
}

func (s *ViewerService) SaveDefaultViewerPage(r DefaultViewerPage) (int64, error) {
	result := s.db.Save(r)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
