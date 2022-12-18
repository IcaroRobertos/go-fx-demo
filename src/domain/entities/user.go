package entities

type User struct {
	ID    int    `json:"id,omitempty" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
