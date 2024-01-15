package client

import (
	"fmt"
	"time"

	ndp "github.com/Lab1038/gondp"
)

func Discover() {
	// 2. Discover

	fmt.Println("Discovering the usable name...")

	// send /BIT/BS/TempName to nearby entity
	// receive "/BIT/AP-1/device01" from nearby entity
	sock, err := ndp.NewNdpSocket()
	if err != nil {
		panic(err)
	}
	defer sock.Close()
	fmt.Println(sock, err)
	buf := make([]byte, 1024)
	fmt.Println("consumer start")
	for i := 0; i < 10; i++ {
		iname := fmt.Sprintf("/gondp/test/%d", i)
		err = sock.SendInterest(iname)
		if err != nil {
			panic(err)
		}
		fmt.Println("consumer send interest: ", iname)

		dname, n, err := sock.ReceiveData(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println("consumer get data", dname, ": ", string(buf[:n]))
		time.Sleep(time.Second)
	}
}
