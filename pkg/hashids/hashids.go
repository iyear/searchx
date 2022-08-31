package hashids

import "github.com/speps/go-hashids/v2"

const salt = "Hp#un&a2YZ3v7S"

func mustGetHash() *hashids.HashID {
	hd := hashids.NewData()
	hd.Salt = salt

	h, err := hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}
	return h
}

func Encode64(ids ...int64) (string, error) {
	return mustGetHash().EncodeInt64(ids)
}

func Encode(ids ...int) (string, error) {
	return mustGetHash().Encode(ids)
}

func Decode(s string) ([]int, error) {
	return mustGetHash().DecodeWithError(s)
}

func Decode64(s string) ([]int64, error) {
	return mustGetHash().DecodeInt64WithError(s)
}
