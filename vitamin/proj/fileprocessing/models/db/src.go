package db

type Source struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	SrcName string `gorm:"type:text;not null"`
}
