package database

import (
	"gorm.io/gorm"
)

// type YesNoBool bool

// func NewNa() (*gorm.DB, error) {
// 	dbc := config.GetConfig().Database
// 	cfg := mysql.Config{
// 		User:      dbc.Username,
// 		Passwd:    dbc.Password,
// 		Net:       dbc.NetworkType,
// 		Addr:      dbc.Address,
// 		DBName:    dbc.Name,
// 		ParseTime: true,
// 		Loc:       time.Local,
// 	}

// 	db, err := gorm.Open(gormMysql.New(gormMysql.Config{
// 		DSNConfig:                 &cfg,
// 		DefaultStringSize:         256,   // default size for string fields
// 		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
// 		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
// 		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
// 		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
// 	}), &gorm.Config{})

// 	// db.AutoMigrate(&remote.Remote{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

func Migrate(db *gorm.DB) {
	// db.AutoMigrate(
	// 	&remote.RemoteJuke{},
	// 	&remote.RemoteOndemand{},
	// 	&remote.RemotePreferences{},
	// 	&remote.PsaSequence{},
	// 	&remote.Remote{},
	// )
}

// func (s *YesNoBool) Scan(value interface{}) error {
// 	b, _ := value.(string)
// 	if b == "y" || b == "Y" {
// 		*s = true
// 	} else {
// 		*s = false
// 	}
// 	return nil
// }

// func (s YesNoBool) Value() (driver.Value, error) {
// 	if s {
// 		return "Y", nil
// 	}
// 	return "N", nil
// }
