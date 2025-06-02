package models

type Note struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	UserID  uint   `gorm:"not null" json:"-"`
}
