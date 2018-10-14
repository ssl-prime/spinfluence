package service

import (
	"fmt"
	"github.com/rightjoin/aqua"
	"spinfluence/api/model"
	"spinfluence/api/util"
	"strconv"
)

type User struct {
	aqua.RestService `prefix:"user/" root:"/" version:"1"`
	signUp           aqua.POST `url:"/signup"`
	login            aqua.POST `url:"/login"`
	details          aqua.GET  `url:"/details"`
	forgotPassword   aqua.POST `url:"/forgot_password"`
	verifyCode       aqua.POST `url:"/verifycode"`
	changePassword   aqua.POST `url:"/change_password/{token}"`
	signOut          aqua.GET  `url:"/signout"`
	getResultUser    aqua.GET  `url:"/getresult_user"`
	takeTest         aqua.POST `url:"/take_test/{id:[0-9]+}/{test_type_id:[0-9]+}"`
	endTest          aqua.POST `url:"/end_test"`
}

//user SignUp
func (usr *User) SignUp(j aqua.Aide) (response interface{}, err error) {
	var (
		reqPayLoad model.User
	)
	fmt.Println("tmp")
	if reqPayLoad, err = util.ValidateSignUp(j); err == nil {
		response, err = util.SignUp(reqPayLoad, j)
	}
	return
}

//user Login
func (usr *User) Login(j aqua.Aide) (response interface{}, err error) {
	var (
		validPWD  bool
		user_name string
	)
	if validPWD, user_name, err = util.ValidateLogin(j); err == nil {
		response, err = util.Login(validPWD, user_name, j)
	}
	return
}

// //user Details
// func (usr *User) Details(j aqua.Aide) (
// 	response interface{}, err error) {
// 	response, err = util.Details(j)
// 	return
// }

//user ForgotPassword
func (usr *User) ForgotPassword(j aqua.Aide) (response interface{}, err error) {

	if err = util.ValidateForgotPassword(j); err == nil {
		response, err = util.ForgotPassword(j)
	}

	return

}

// //user VerifyCode
// func (usr *User) VerifyCode(j aqua.Aide) (
// 	response interface{}, err error) {
// 	response, err = util.VerifyCode(j)
// 	return
// }

//user  ChangePassword
func (usr *User) ChangePassword(token string, j aqua.Aide) (response interface{}, err error) {
	var (
		tokenStaus bool
		PWDPload   model.ChangePWD
	)
	if tokenStaus, PWDPload, err = util.ValidateChangePassword(token, j); err == nil {
		if tokenStaus {
			response, err = util.ChangePassword(PWDPload, j)
		}
	}
	return
}

//user SignOut
func (usr *User) SignOut(j aqua.Aide) (response interface{}, err error) {
	var (
		validSession bool
	)
	if validSession, err = util.ValidateSignOut(j); err == nil {
		if validSession {
			response, err = util.SignOut(validSession, j)
		}
	}
	return
}

func (usr *User) GetResultUser(j aqua.Aide) (response interface{}, err error) {
	j.LoadVars()
	var (
		ID    int
		valid bool
	)
	if ID, err = strconv.Atoi(j.Request.Header.Get("fkUserID")); err == nil {
		if valid, err = util.ValidateGetResultUser(ID, j); err == nil && valid {
			response, err = util.GetResultUser(ID, j)
		}
	}
	return
}

//TakeTest
func (usr *User) TakeTest(id string, test_type_id int, j aqua.Aide) (response interface{}, err error) {
	var (
		validResponse bool
	)
	if validResponse, err = util.ValidateTakeTest(id, test_type_id, j); err == nil {
		if validResponse {
			response, err = util.TakeTest(id, test_type_id, j)
		}
	}
	return
}

//EndTest
func (usr *User) EndTest(j aqua.Aide) (response interface{}, err error) {
	var (
		validResponse bool
		paylaod       model.TestResponse
	)
	if validResponse, paylaod, err = util.ValidateEndTest(j); err == nil {
		if validResponse {
			response, err = util.EndTest(paylaod, j)
		}
	}
	return
}
