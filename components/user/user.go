package user

import "gorm.io/gorm"

//User - Model
type User struct {
	gorm.Model
	Name       string `json:"name"`
	Email      string `json:"email"`
	FirebaseID string `json:"firebase_id"`
}
