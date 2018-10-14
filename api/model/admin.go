package model

type Result struct {
	UserID     int     `json:"fk_user_id" gorm:"column:fk_user_id"`
	Score      float64 `json:"score" gorm:"column:score"`
	TestTypeID int     `json:"fk_test_type_id" gorm:"column:fk_test_type_id"`
	Attempt    int     `json:"attempt" gorm:"column:attempt"`
}
