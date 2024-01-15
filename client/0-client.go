package client

import (
	"fmt"
)

func Bootstrap() {
	fmt.Println("This is client bootstrapping program.")

	// 1. Init
	InitClient()

	// 2. Discover
	Discover()

	// 3. Qualifiaction
	Qualifiaction()

}
