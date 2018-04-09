package utils

import (
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
	graphql "github.com/graph-gophers/graphql-go"
)

// EncodeUint64 returns an 8-byte big endian representation of uint64 i.
func EncodeUint64(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

// DecodeUint64 returns a uint64 representation of 8-byte big endian b.
func DecodeUint64(b []byte) uint64 {
	var u uint64
	u = binary.BigEndian.Uint64(b)
	return u
}

// FetchID gets an ID from a boltdb bucket
func FetchID(id graphql.ID, b *bolt.Bucket) []byte {
	i, _ := strconv.ParseInt(string(id), 10, 64)
	ib := EncodeUint64(uint64(i))
	return b.Get(ib)
}

// IDToUintBytes takes a grapql.ID integer, and converts it to a Big Endian version of a Uint64
func IDToUintBytes(id graphql.ID) []byte {
	i, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return EncodeUint64(i)
}
