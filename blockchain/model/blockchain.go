package model

import (
	"spinfluence/api/model"
)

//block
type Block struct {
	Index     int
	Timestamp string
	Answer    model.TestResponse
	Hash      string
	PrevHash  string
}

//message
type Message struct {
	BPM int
}
