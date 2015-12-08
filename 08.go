package main
import (
	"fmt"
	"strconv"
	"os"

	"github.com/ifross89/adventofcode/getinput"
)



func main() {


	diffA := 0
	diffB := 0
	for _, literalStr := range getinput.MustGet(8, os.Getenv("ADVENT_SESSION"))  {
		quotedStr := strconv.Quote(literalStr)
		memoryStr, _ := strconv.Unquote(literalStr)
		diffA += len(literalStr) - len(memoryStr)
		diffB += len(quotedStr) - len(literalStr)
	}
	fmt.Println("literal length - in memory length =", diffA)
	fmt.Println("quoted length - literal length =", diffB)
}
