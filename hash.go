package main

import (
	"encoding/hex"
	"hash/fnv"
)

func getHash(s string) string {
	hasher := fnv.New32a()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}
