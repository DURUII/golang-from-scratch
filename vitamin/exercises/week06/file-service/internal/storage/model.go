package storage

type Item struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	UUID     string `gorm:"unique"`
	Type     string `gorm:"type:text"`
	FilePath string `gorm:"type:text"`
	FileSize int64  `gorm:"type:int"`
	IsDelete int    `gorm:"default:0"`
}
