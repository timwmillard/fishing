package fishing

import (
	"encoding/json"
	"errors"

	"github.com/speps/go-hashids/v2"
)

const (
	hashMinLength = 10
	hashAlphabet  = "abcdefghijklmnopqrstuvwxyz1234567890"
	hashSalt      = "this is my secret"
)

type HashID int

func (id HashID) Hash() (string, error) {
	hash, err := hashids.NewWithData(hashData())
	if err != nil {
		return "", err
	}

	str, err := hash.Encode([]int{int(id)})
	if err != nil {
		return "", err
	}
	return str, nil
}

func NewHashID(h string) (HashID, error) {
	hash, err := hashids.NewWithData(hashData())
	if err != nil {
		return 0, err
	}
	// str := hash.Decode(string(h)) // TODO: deprecated
	str, err := hash.DecodeWithError(h)
	if err != nil {
		return 0, err
	}
	if len(str) == 0 {
		return 0, errors.New("invalid hash")
	}
	return HashID(str[0]), nil
}

func (id HashID) MarshalJSON() ([]byte, error) {
	h, err := id.Hash()
	if err != nil {
		return nil, err
	}
	return json.Marshal(h)
}

func (id *HashID) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	h, err := NewHashID(j)
	if err != nil {
		return err
	}
	*id = h
	return nil
}

func hashData() *hashids.HashIDData {
	return &hashids.HashIDData{
		Alphabet:  hashAlphabet,
		MinLength: hashMinLength,
		Salt:      hashSalt,
	}
}
