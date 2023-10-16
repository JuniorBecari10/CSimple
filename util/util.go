package util

import (
	"fmt"
  "strconv"

	"csimple/token"
)

type FileData struct {
	Name  string
	Lines []string
}

/*
Error at 3:1 | Invalid Number: '4.5.3'
  --> src/main.sm
 3 | 4.5.3
     ^

Error ...
*/

func ThrowError(data FileData, pos token.Position, msg string) {
	fmt.Printf("Error at %d:%d | %s\n", pos.Line+1, pos.Col+1, msg)
	fmt.Printf("  --> %s\n", data.Name)
	fmt.Printf(" %d | %s\n", pos.Line+1, data.Lines[pos.Line])

	for i := 0; i < pos.Col + len(strconv.Itoa(pos.Col)) + 4; i++ {
		fmt.Print(" ")
	}

	fmt.Println("^")
}
