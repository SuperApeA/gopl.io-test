package main

////练习 8.2： 实现一个并发FTP服务器。服务器应该解析客户端来的一些命令，比如cd命令来切换目录，ls来列出目录内文件，get和send来传输文件，close来关闭连接。你可以用标准的ftp命令来作为客户端，或者也可以自己实现一个。
//
//
//
//import (
//	"log"
//	"net"
//)
//
//func main() {
//	conn, err := net.Dial("localhost", "8888")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//}
