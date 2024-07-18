package service

type User struct {
	ID    string `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
