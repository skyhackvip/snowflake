package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	u, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
}
