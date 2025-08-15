package db

type Word struct {
	ID           uint          `gorm:"primaryKey;autoIncrement"`
	SrcID        uint          // 外键关联 SrcID
	Word         string        `gorm:"type:text;not null"`
	IsDelete     int           `gorm:"default:0"`
	Translations []Translation `gorm:"foreignKey:WordID"`
	Phrases      []Phrase      `gorm:"foreignKey:WordID"`
}
