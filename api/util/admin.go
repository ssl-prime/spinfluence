package util

import (
	//"encoding/json"
	"fmt"
	//"github.com/asaskevich/govalidator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rightjoin/aqua"
	//"net/http"
	"spinfluence/api/model"
	// crypt "spinfluence/encrypt_decrypt"
	// nt "spinfluence/notification/email/util"
	// "time"
	"strconv"
)

//ValiadateGetResult
func ValidateGetResult(j aqua.Aide) (err error) {
	j.LoadVars()
	var fkUserID int
	if fkUserID, err = strconv.Atoi(j.Request.Header.Get("fkUserID")); err == nil {
		err = Admin(fkUserID)
	}
	return
}

//GetResult
func GetResult(j aqua.Aide) (response interface{}, err error) {
	var (
		Conn    *gorm.DB
		results []model.Result
	)

	if Conn, err = dbConn(); err == nil {
		result := `Select fk_user_id, score, fk_test_type_id, attempt from result;`
		if err = Conn.Raw(result).Find(&results).Error; err == nil {
			response = results
		}
	}
	return
}

//ValiadateGetResultByID
func ValidateGetResultByID(ID int, j aqua.Aide) (err error) {
	j.LoadVars()
	var fkUserID int
	if fkUserID, err = strconv.Atoi(j.Request.Header.Get("fkUserID")); err == nil {
		if err = Admin(fkUserID); err == nil {
			err = User(ID)
		}
	}
	return
}

//GetResultByID
func GetResultByID(ID int, j aqua.Aide) (response interface{}, err error) {
	var (
		Conn    *gorm.DB
		results []model.Result
	)
	if Conn, err = dbConn(); err == nil {
		result := `Select fk_user_id, score, fk_test_type_id, attempt from result where fk_user_id= ?;`
		if err = Conn.Raw(result, ID).Find(&results).Error; err == nil {
			response = results
		}
	}
	return
}

//ValiadateGetUserByID
func ValidateGetUser(j aqua.Aide) (err error) {
	j.LoadVars()
	var fkUserID int
	if fkUserID, err = strconv.Atoi(j.Request.Header.Get("fkUserID")); err == nil {
		err = Admin(fkUserID)
	}
	return
}

//GetUserByID
func GetUser(j aqua.Aide) (response interface{}, err error) {
	var (
		Conn *gorm.DB
		usr  []model.User
	)
	if Conn, err = dbConn(); err == nil {
		getUser := `Select user_id, user_name, first_name, last_name, email_id,
		 contact_number	from user where user_type = ?;`
		if err = Conn.Raw(getUser, "user").Find(&usr).Error; err == nil {
			response = usr
		}
	}
	return
}

//ValiadateGetUserByID
func ValidateGetUserByID(ID int, j aqua.Aide) (err error) {
	j.LoadVars()
	var fkUserID int
	if fkUserID, err = strconv.Atoi(j.Request.Header.Get("fkUserID")); err == nil {
		if err = Admin(fkUserID); err == nil {
			err = User(ID)
		}
	}
	return
}

//GetUserByID
func GetUserByID(ID int, j aqua.Aide) (response interface{}, err error) {
	var (
		Conn *gorm.DB
		usr  model.User
	)
	if Conn, err = dbConn(); err == nil {
		getUser := `Select user_id, user_name, first_name, last_name, email_id, contact_number
					from user where user_type = ? AND user_id = ?;`
		if err = Conn.Raw(getUser, "user", ID).Find(&usr).Error; err == nil {
			response = usr
		}
	}
	return
}

// validateAdmin
func Admin(fkUserID int) (err error) {
	var (
		Conn     *gorm.DB
		usr_name model.User
	)
	if Conn, err = dbConn(); err == nil {
		validateAdmin := `Select user_name from user where user_id = ? AND user_type = ?;`
		if err = Conn.Raw(validateAdmin, fkUserID, "Admin").Find(&usr_name).Error; err == nil {
			fmt.Println(usr_name.UserName)
		}
	}
	return
}

//validateUser
func User(ID int) (err error) {
	var (
		Conn     *gorm.DB
		usr_name model.User
	)
	if Conn, err = dbConn(); err == nil {
		validateAdmin := `Select user_name from user where user_id = ? AND user_type = ?;`
		if err = Conn.Raw(validateAdmin, ID, "User").Find(&usr_name).Error; err == nil {
			fmt.Println(usr_name.UserName)
		}
	}
	return
}
