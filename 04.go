package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	i := int64(0)

	// Add secret key in here (extra padding to stop reallocations)
	in := []byte{'i', 'w', 'r', 'u', 'p', 'v', 'q', 'b', 0, 0, 0, 0, 0, 0, 0, 0}
	var result1 int64
	var result2 int64
	var finished bool
	for {
		text := strconv.AppendInt(in[:8], i, 10)
		hash := md5.Sum(text)
		if hash[0] == 0 && hash[1] == 0 && hash[2] < 16 {

			if result1 == 0 {
				result1 = i
				finished = true
			}
			if hash[2] == 0 {
				if result2 == 0 {
					result2 = i
				}
				if finished {
					break
				}
			}
		}
		i++
	}
	fmt.Println("5 zero result: ", result1)
	fmt.Println("6 zero result: ", result2)
}
