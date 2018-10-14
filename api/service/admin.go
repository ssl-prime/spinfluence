package service

import (
	"github.com/rightjoin/aqua"
	//"spinfluence/api/model"
	"spinfluence/api/util"
)

type Admin struct {
	aqua.RestService `prefix:"admin/" root:"/" version:"1"`
	getResult        aqua.POST `url:"/getresult"`
	getResultByID    aqua.POST `url:"/getresult/{ID:[0-9]+}"`
	getUser          aqua.POST `url:"/getuser"`
	getUserByID      aqua.POST `url:"/getuser/{ID:[0-9]+}"`
}

//GetResult
func (adm *Admin) GetResult(j aqua.Aide) (response interface{}, err error) {
	if err = util.ValidateGetResult(j); err == nil {
		response, err = util.GetResult(j)
	}
	return
}

//GetResultByID
func (adm *Admin) GetResultByID(ID int, j aqua.Aide) (response interface{}, err error) {
	if err = util.ValidateGetResultByID(ID, j); err == nil {
		response, err = util.GetResultByID(ID, j)
	}
	return
}

//GetUser
func (adm *Admin) GetUser(j aqua.Aide) (response interface{}, err error) {
	if err = util.ValidateGetUser(j); err == nil {
		response, err = util.GetUser(j)
	}
	return
}

//GetUserByID
func (adm *Admin) GetUserByID(ID int, j aqua.Aide) (response interface{}, err error) {
	if err = util.ValidateGetUserByID(ID, j); err == nil {
		response, err = util.GetUserByID(ID, j)
	}
	return
}
