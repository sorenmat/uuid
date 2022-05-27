package main

import (
	"fmt"
	"github.com/google/uuid"
	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	u := uuid.NewString()
	fmt.Println(u)
	clipboard.Write(clipboard.FmtText, []byte(u))

}
