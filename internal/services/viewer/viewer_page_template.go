package viewer

import "go.mongodb.org/mongo-driver/bson/primitive"

type ViewerPageTemplate struct {
	Id   *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string              `json:"name" bson:"name"`
	Html string              `json:"html" bson:"html"`
}

// type RemoteViewerPageTemplate struct {
// 	Id       int64  `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
// 	Name     string `json:"name" gorm:"column:name"`
// 	Html     string `json:"html" gorm:"column:html"`
// 	IsActive bool   `json:"isActive" gorm:"column:isActive"`
// }

// func (RemoteViewerPageTemplate) TableName() string {
// 	return "REMOTE_VIEWER_PAGE_TEMPLATE"
// }

// func (s *ViewerService) GetViewerPageTemplate(id int64) (*RemoteViewerPageTemplate, error) {
// 	var r RemoteViewerPageTemplate
// 	result := s.db.First(&r, "id = ?", id)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return &r, nil
// }

// func (s *ViewerService) SaveViewerPageTemplate(r RemoteViewerPageTemplate) (int64, error) {
// 	result := s.db.Save(r)
// 	if result.Error != nil {
// 		return 0, result.Error
// 	}
// 	return result.RowsAffected, nil
// }

// func (s *ViewerService) DeleteViewerPageTemplate(r RemoteViewerPageTemplate) (int64, error) {
// 	result := s.db.Delete(r)
// 	if result.Error != nil {
// 		return 0, result.Error
// 	}
// 	return result.RowsAffected, nil
// }
