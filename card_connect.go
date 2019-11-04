// connects to card scanner & does reads

package main

import (
	"fmt"
	"os"

	"github.com/tarm/serial"
)

func main() {
	c0 := &serial.Config{
		Name:        "USB\\VID_0ACD&PID_0520\\8&1b6f5f00&0&3",
		Baud:        9600,
		ReadTimeout: 1,
		Size:        8,
	}
	stream, err := serial.OpenPort(c0)
	if err != nil {
		fmt.Println("Error EstablishContext:", err)
		return
	}
	defer stream.Close()
	fmt.Printf("%s", stream)

	buf := make([]byte, 1024)
	for {
		n, err := stream.Read(buf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		s := string(buf[:n])
		fmt.Println(s)
	}

}
