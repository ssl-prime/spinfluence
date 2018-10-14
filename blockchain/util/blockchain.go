package util

import (
	"crypto/sha256"
	"encoding/hex"
	mdl "spinfluence/api/model"
	"spinfluence/blockchain/model"
	//"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/rightjoin/aqua"
	"sync"
	"time"
)

//golbal blockchain
var Blockchain []model.Block

//lock
var mutex = &sync.Mutex{}

//validateGetBlock
func ValidateGetBlock(j aqua.Aide) (err error) {

	return nil
}

//GetBlock
func GetBlock(j aqua.Aide) (response interface{}, err error) {

	return Blockchain, nil
}

//ValidateAddBlock
func ValidateAddBlock(j aqua.Aide) (err error) {

	return nil
}

//Todo we have make queue for multiple request
//AddBlock
func AddBlock(ans mdl.TestResponse) (response interface{}, err error) {

	if len(Blockchain) == 0 {
		go func() {
			t := time.Now()
			genesisBlock := model.Block{}
			genesisBlock = model.Block{0, t.String(), ans, calculateHash(genesisBlock), ""}
			spew.Dump(genesisBlock)

			mutex.Lock()
			Blockchain = append(Blockchain, genesisBlock)
			mutex.Unlock()
		}()
	} else {
		mutex.Lock()
		var newBlock model.Block
		newBlock, err = generateBlock(Blockchain[len(Blockchain)-1], ans)
		mutex.Unlock()

		if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
			Blockchain = append(Blockchain, newBlock)
			spew.Dump(Blockchain)
		}

	}

	return "done", err
}

//calculateHash
func calculateHash(block model.Block) string {

	record := string(block.Index) + block.Timestamp + string(block.Answer.UserID) + block.PrevHash

	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

//generateBlock
func generateBlock(oldBlock model.Block, ans mdl.TestResponse) (model.Block, error) {

	var newBlock model.Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Answer = ans
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

//isBlockValid
func isBlockValid(newBlock model.Block, oldBlock model.Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

//replaceChain
func replaceChain(newBlocks []model.Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
