package viewer

import "go.mongodb.org/mongo-driver/bson/primitive"

type ViewerPage struct {
	Id       *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string              `json:"name" bson:"name"`
	Html     string              `json:"html" bson:"html"`
	Title    string              `json:"title" bson:"title"`
	IconLink string              `json:"iconLink" bson:"iconLink"`
}

// type RemoteViewerPage struct {
// 	Id               int64  `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
// 	RemoteId         int64  `json:"remoteId" gorm:"column:remoteId"`
// 	ViewerPageName   string `json:"viewerPageName" gorm:"column:viewerPageName"`
// 	ViewerPageActive bool   `json:"viewerPageActive" gorm:"column:viewerPageActive"`
// 	ViewerPageHtml   string `json:"viewerPageHtml" gorm:"column:viewerPageHtml"`
// }

// func (RemoteViewerPage) TableName() string {
// 	return "REMOTE_VIEWER_PAGE"
// }

// func (s *ViewerService) GetViewerPage(id int64) (*RemoteViewerPage, error) {
// 	var r RemoteViewerPage
// 	result := s.db.First(&r, "remoteId = ?", id)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &r, nil
// }

// func (s *ViewerService) SaveViewerPage(r RemoteViewerPage) (int64, error) {
// 	result := s.db.Save(r)
// 	if result.Error != nil {
// 		return 0, result.Error
// 	}
// 	return result.RowsAffected, nil
// }
