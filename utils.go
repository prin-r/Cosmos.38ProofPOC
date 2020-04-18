package main

import (
	"encoding/base64"
	"encoding/hex"
	"reflect"
	"time"

	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

func Base64ToBytes(s string) []byte {
	b64, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err.Error())
	}
	return b64
}

func EncodeKV(key string, value []byte) []byte {
	lenKey := uint8(len([]byte(key)))
	lenValue := uint8(len(value))
	return append(append([]byte{lenKey}, []byte(key)...), append([]byte{lenValue}, value...)...)
}

func ParseTime(s string) time.Time {
	t, _ := time.Parse(time.RFC3339, s)
	return t
}

func HexToBytes(s string) []byte {
	x, _ := hex.DecodeString(s)
	return x
}

func IsTypedNil(o interface{}) bool {
	rv := reflect.ValueOf(o)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

func EncodeIfNotNil(item interface{}) []byte {
	if item == nil || IsTypedNil(item) {
		return nil
	}
	return (amino.NewCodec()).MustMarshalBinaryBare(item)
}

func LeafHash(item []byte) []byte {
	// leaf prefix is 0
	return tmhash.Sum(append([]byte{0}, item...))
}

func BranchHash(left, right []byte) []byte {
	// branch prefix is 1
	return tmhash.Sum(append([]byte{1}, append(left, right...)...))
}

func HashHeader(items [][]byte) []byte {
	if len(items) != 14 {
		panic("hashHeader error, len of items must be 14")
	}
	return BranchHash(
		BranchHash(
			BranchHash(
				BranchHash(LeafHash(items[0]), LeafHash(items[1])),
				BranchHash(LeafHash(items[2]), LeafHash(items[3])),
			),
			BranchHash(
				BranchHash(LeafHash(items[4]), LeafHash(items[5])),
				BranchHash(LeafHash(items[6]), LeafHash(items[7])),
			),
		),
		BranchHash(
			BranchHash(
				BranchHash(LeafHash(items[8]), LeafHash(items[9])),
				BranchHash(LeafHash(items[10]), LeafHash(items[11])),
			),
			BranchHash(LeafHash(items[12]), LeafHash(items[13])),
		),
	)
}
