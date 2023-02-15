package file

import (
	"fmt"
	"os"
)

func ReadFile() {
	if len(os.Args) < 2 {
		panic("No file to read")
	}
	f := os.Args[1]
	if _, err := os.Stat(f); err != nil {
		panic(err)
	}
	file, err := os.ReadFile(f)
	for i, v := range file {
		if v == 32 {
			fmt.Println(i, string(v)
		}
	}
	// convert file to entity.author
	//var a entity.Author

	if err != nil {
		panic(err)
	}
	println(string(file))
}
