package model

type Result struct {
	UserID     int     `json:"fk_user_id" gorm:"column:fk_user_id"`
	Score      float64 `json:"score" gorm:"column:score"`
	TestTypeID int     `json:"fk_test_type_id" gorm:"column:fk_test_type_id"`
	Attempt    int     `json:"attempt" gorm:"column:attempt"`
}

//
type UserDetails struct {
	UserID        int    `json:"user_id" `
	FirstName     string `json:"first_name" valid:"required"`
	LastName      string `json:"last_name"  valid:"required"`
	UserName      string `json:"user_name"  valid:"required"`
	EmailID       string `json:"email_id"   valid:"required"`
	ContactNumber string `json:"contact_number" valid:"required"`
	Password      string `json:"-"   valid:"required"`
	Sex           string `json:"-" valid:"required"`
}
