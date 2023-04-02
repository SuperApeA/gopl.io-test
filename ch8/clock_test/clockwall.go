package main

import (
	"fmt"
	"net"
)

func main() {
	for port, _ := range portToArea {
		connToClock, err := net.Dial("tcp", "localhost:"+port)
		if err != nil {
			fmt.Println("Error connecting:", err)
			return
		}
		// 读取返回的时间
		time := make([]byte, 128)
		_, err = connToClock.Read(time)
		if err != nil {
			fmt.Println("Error connect read:", err)
			return
		}
		fmt.Println(string(time))
		if err := connToClock.Close(); err != nil {
			fmt.Println("Error close connect:", err)
			return
		}
	}
}
