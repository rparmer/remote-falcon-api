package remote

// const collection = "remote_juke"

type RemoteOndemand struct {
	Id       int64  `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
	RemoteId int64  `json:"remoteId" gorm:"column:remoteId"`
	Playlist string `json:"playlist" gorm:"column:playlist"`
}

func (RemoteOndemand) TableName() string {
	return "REMOTE_ONDEMAND"
}

// func (s *RemoteService) GetOnDemand(id int64) (*RemoteOndemand, error) {
// 	var r RemoteOndemand
// 	result := s.db.First(&r, "remoteId = ?", id)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &r, nil
// }

// func (s *RemoteService) SaveOnDemand(r RemoteOndemand) (int64, error) {
// 	result := s.db.Save(r)
// 	if result.Error != nil {
// 		return 0, result.Error
// 	}
// 	return result.RowsAffected, nil
// }
