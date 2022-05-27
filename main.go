package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	u := uuid.NewString()
	fmt.Println(u)
}
