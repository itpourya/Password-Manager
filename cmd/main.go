package main

import (
	"fmt"

	"github.com/nothyphen/Password-Manager/routes"
)

func main() {
	route := routes.Urls()
	err := route.Run()
	if err != nil{
		fmt.Println(err)
	}
}