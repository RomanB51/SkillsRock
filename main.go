package main

import (
	//"github.com/gofiber/fiber/v2"

	"fmt"
)

func main() {
	var err error
	if err = Conn_to_DB(); err != nil {
		fmt.Println(err)
	}

}
