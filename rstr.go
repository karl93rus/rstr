package rstr

import (
	"math/rand"
	"time"
)

var symSet []byte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_")
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func Genstr(l int) string {
	resKey := []byte{}

	for i := 0; i < l; i++ {
		resKey = append(resKey, symSet[rnd.Int63n(int64(len(symSet)))])
	}
	return string(resKey)
}
