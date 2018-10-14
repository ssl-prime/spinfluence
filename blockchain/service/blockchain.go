package service

import (
	//"blockchain/api/model"
	"fmt"
	"github.com/rightjoin/aqua"
	"spinfluence/blockchain/util"
)

//blockchain api
type BlockChain struct {
	aqua.RestService `prefix:"blockchain/" root:"/" version:"1"`

	getBlock aqua.GET `url:"/get_block"`
	//addBlock aqua.POST `url:"/add_block"`
}

//getBlocks of chain
func (blc *BlockChain) GetBlock(j aqua.Aide) (
	response interface{}, err error) {
	fmt.Println("get")
	if err = util.ValidateGetBlock(j); err == nil {
		response, err = util.GetBlock(j)
	}
	return
}

// //addBlock
// func (blc *BlockChain) AddBlock(j aqua.Aide) (
// 	response interface{}, err error) {
// 	fmt.Println("add")
// 	if err = util.ValidateAddBlock(j); err == nil {
// 		response, err = util.AddBlock(j)
// 	}
// 	return
// }
