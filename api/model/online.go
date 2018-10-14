package model

import (
//"time"
)

//user model
type User struct {
	FirstName string `json:"first_name" valid:"required"`
	LastName  string `json:"last_name"  valid:"required"`
	UserName  string `json:"user_name"  valid:"required"`
	EmailID   string `json:"email_id"   valid:"required"`
	ContactNO string `json:"contact_number"  valid:"required"`
	Password  string `json:"password"   valid:"required"`
	Sex       string `json:"sex" valid:"required"`
	Country   string `json:"country" valid:"required"`
	IsActive  int    `json:"-"`
}

//login
type LoginInfo struct {
	UserName string `json:"user_name" `
	EmailID  string `json:"email_id" `
	Password string `json:"password" valid:"required"`
}

//ForgotPassword
type ForgotPassword struct {
	UserName string `json:"user_name" valid:"required" gorm:"column:user_name"`
	EmailID  string `json:"email_id" valid:"required" gorm:"column:email_id"`
	Password string `json:"password" gorm:"column:password"`
}

//
type PassWord struct {
	PassWord string `json:"password" gorm:"column:password"`
}

//change password info
type ChangePWD struct {
	UserName   string `json:"user_name" gorm:"column:user_name"`
	Password   string `json:"password" gorm:"column:password"`
	ExpiryTime string `json:"expiry_time" gorm:"column:expiry_time"`
}

// answerModel
type Answer struct {
	ID       int    `json:"id"`
	Response string `json:"response"`
}

//TestResponse
type TestResponse struct {
	UserID         int      `json:"user_id" valid:"required"`
	UserName       string   `json:"user_name"`
	TotalNoQst     int      `json:"total_qst"`
	TotalQstAttmpt int      `json:"total_qst_attmpt"`
	TotalRes       []Answer `json:"total_res"`
	TestTypeID     int      `json:"test_type_id" valid:"required"`
	AttemptNo      int      `json:"attempt_no"`
}
