package main

import (
	"app/cmd"
	"fmt"
	"os"
)

func main() {
	if err := cmd.NewCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}