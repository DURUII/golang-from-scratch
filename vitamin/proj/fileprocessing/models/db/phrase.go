package db

type Phrase struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	WordID      uint   // 外键关联 Word.ID
	Phrase      string `gorm:"type:text"`
	Translation string `gorm:"type:text"`
	IsDelete    int    `gorm:"default:0"`
}
