package main

import (
	"fmt"
	"github.com/shxsun/filedist/fire/utils"
)

func main() {
	tidx := utils.NewTruncIndex()
	err := tidx.Add("hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	id := tidx.New()
	fmt.Println(id)
	fmt.Println(tidx.Get("13"))
	fmt.Println(tidx.Get("hs"))
}
