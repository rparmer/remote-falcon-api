package remote

import (
	"time"
)

type Remote struct {
	Id            int       `json:"id" bson:"_id,omitempty"`
	RemoteToken   string    `json:"remoteToken" gorm:"column:remoteToken"`
	RemoteUrl     string    `json:"remoteUrl" gorm:"column:remoteUrl"`
	RemoteName    string    `json:"remoteName" gorm:"column:remoteName"`
	Username      string    `json:"username" gorm:"column:username"`
	Password      string    `json:"password" gorm:"column:password"`
	CreatedDate   time.Time `json:"createdDate" gorm:"column:createdDate;autoCreateTime"`
	LastLoginDate time.Time `json:"lastLoginDate" gorm:"column:lastLoginDate"`
	LastLoginIp   string    `json:"lastLoginIp" gorm:"column:lastLoginIp"`
}

func (Remote) TableName() string {
	return "REMOTE"
}

// func (s *RemoteService) GetRemote(id int64) (*Remote, error) {
// 	var r Remote
// 	result := s.db.First(&r, "id = ?", id)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &r, nil
// }

// func (s *RemoteService) GetRemoteByName(name string) (*Remote, error) {
// 	var r Remote
// 	result := s.db.First(&r, "remoteName = ?", name)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &r, nil
// }

// func (s *RemoteService) SaveRemote(r Remote) (int64, error) {
// 	result := s.db.Save(r)
// 	if result.Error != nil {
// 		return 0, result.Error
// 	}
// 	return result.RowsAffected, nil
// }
