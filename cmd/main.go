package main

import (
	"fmt"
	"reader-csv/pkg/api"
	"reader-csv/pkg/file"
)

func main() {
	fmt.Println("Start Reader CSV")
	file.ReadFile()
	api.Start()
}
