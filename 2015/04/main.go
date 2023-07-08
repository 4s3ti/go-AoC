package main

import (
	"crypto/md5"
	"fmt"
	"strings"

)

const (
	input = "bgvyzdsv"
)

func zeros(n int) string {
	return strings.Repeat("0", n)
}

func isZeros(hashStart string, z int) bool {
	if hashStart ==	zeros(z) {
		return true
	}
	return false
}


func findHash(key string, z int) (int, string) {
	int := 0
	var md5hash string
	for {
		toHash := key + fmt.Sprint(int)
		data := []byte(toHash)
		md5hash = fmt.Sprintf("%x", md5.Sum(data))
		hashStart := md5hash[0:z]
		if isZeros(hashStart, z) {
			break
		}
		int++
	}
	return int, md5hash
}

func Part1(key string) {
	fmt.Println(findHash(key, 5))
}

func Part2(key string) {
	fmt.Println(findHash(key, 6))
}

func main() {
	Part1(input)
	Part2(input)
}
