/*
 * Copyright (c) 2018 QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package util

import (
	"encoding/hex"

	"golang.org/x/crypto/blake2b"
)

func ReverseBytes(str []byte) (result []byte) {
	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, str[i])
	}
	return result
}

func HexToBytes(s string) []byte {
	bytes, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return bytes
}

func Hex32ToBytes(s string) [32]byte {
	var res [32]byte
	bytes := HexToBytes(s)
	copy(res[:], bytes)
	return res
}

func Hex64ToBytes(s string) [64]byte {
	var res [64]byte
	bytes := HexToBytes(s)
	copy(res[:], bytes)
	return res
}

func Hash256(data ...[]byte) []byte {
	d, _ := blake2b.New256(nil)
	for _, item := range data {
		d.Write(item)
	}
	return d.Sum(nil)
}

func Hash(size int, data ...[]byte) []byte {
	d, _ := blake2b.New(size, nil)
	for _, item := range data {
		d.Write(item)
	}
	return d.Sum(nil)
}

// TrimQuotes trim quotes of string if quotes exist
func TrimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}
