package fishing

import (
	"github.com/speps/go-hashids/v2"
)

var HashSalt = "this is my secret"

type HashID string

func NewHashID(id int) HashID {
	data := hashids.NewData()
	data.Salt = HashSalt
	data.MinLength = 10
	hash, err := hashids.NewWithData(data)
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
	data := hashids.NewData()
	data.Salt = HashSalt
	data.MinLength = 10
	hash, err := hashids.NewWithData(data)
	if err != nil {
		panic(err)
	}
	str := hash.Decode(string(h))
	if len(str) == 0 {
		return 0
	}
	return str[0]
}
