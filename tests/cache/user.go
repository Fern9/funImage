package main

import (
	"github.com/zhusun/funImage/pkg/cache/user"
	"fmt"
)

func main() {
	token := "ded"
	u := user.Get(token)
	fmt.Print(u)
}