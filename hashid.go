package fishing

import (
	"github.com/speps/go-hashids/v2"
)

const (
	hashMinLength = 10
	hashAlphabet  = "abcdefghijklmnopqrstuvwxyz1234567890"
	hashSalt      = "this is my secret"
)

type HashID string

func NewHashID(id int) HashID {
	hash, err := hashids.NewWithData(hashData())
	if err != nil {
		panic(err)
	}

	str, err := hash.Encode([]int{id})
	if err != nil {
		panic(err)
	}
	return HashID(str)
}

func (h HashID) ID() int {
	hash, err := hashids.NewWithData(hashData())
	if err != nil {
		panic(err)
	}
	str := hash.Decode(string(h))
	if len(str) == 0 {
		return 0
	}
	return str[0]
}

func hashData() *hashids.HashIDData {
	return &hashids.HashIDData{
		Alphabet:  hashAlphabet,
		MinLength: hashMinLength,
		Salt:      hashSalt,
	}
}
