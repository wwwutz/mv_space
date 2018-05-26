package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	old := os.Args[1]
	wrd := strings.Split(strings.ToLower(old), "_")
	for i := range wrd {
		wrd[i] = strings.Title(wrd[i])
	}
	new := strings.Join(wrd, " ")

	//	fmt.Println(old,new)

	err := os.Rename(old, new)
	if err != nil {
		fmt.Println(err)
		return
	}
}
