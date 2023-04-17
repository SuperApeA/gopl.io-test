package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("please input ip and ardress like '127.0.0.1:8888'")
	}
	ipAddress := os.Args[1]
	conn, err := net.Dial("tcp", ipAddress)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	// 循环读取用户输入并发送给FTP服务器
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// 发送命令到FTP服务器
		fmt.Fprintf(conn, "%s\n", input)

		// 处理服务器响应
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error receiving response:", err)
			return
		}
		fmt.Println(response)

		// 如果是close命令，则退出循环并关闭连接
		if input == "close" {
			break
		}
	}
}
