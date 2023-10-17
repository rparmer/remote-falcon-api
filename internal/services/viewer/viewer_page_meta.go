package viewer

type ViewerPageMeta struct {
	ViewerPageMetaKey  int64  `json:"viewerPageMetaKey" gorm:"column:viewerPageMetaKey;autoIncrement;primaryKey"`
	RemoteToken        string `json:"remoteToken" gorm:"column:remoteToken"`
	ViewerPageTitle    string `json:"viewerPageTitle" gorm:"column:viewerPageTitle"`
	ViewerPageIconLink string `json:"viewerPageIconLink" gorm:"column:viewerPageIconLink"`
}

func (ViewerPageMeta) TableName() string {
	return "VIEWER_PAGE_META"
}

func (s *ViewerService) GetViewerPageMeta(id int64) (*ViewerPageMeta, error) {
	var v ViewerPageMeta
	result := s.db.First(&v, "remoteId = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &v, nil
}

func (s *ViewerService) SaveViewerPageMeta(r ViewerPageMeta) (int64, error) {
	result := s.db.Save(r)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
