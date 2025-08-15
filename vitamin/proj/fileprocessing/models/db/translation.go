package db

type Translation struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	WordID      uint   // 外键关联 Word.ID
	Translation string `gorm:"type:text"`
	Type        string `gorm:"type:text"`
	IsDelete    int    `gorm:"default:0"`
}
