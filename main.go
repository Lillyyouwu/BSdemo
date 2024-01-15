package main

import (
	trustanchor "bootstrapper/TrustAnchor"
	"bootstrapper/client"
	"fmt"
)

func main() {
	fmt.Println("This is a simple bootstrapper.")
	client.Bootstrap()
	trustanchor.Listen()
}
