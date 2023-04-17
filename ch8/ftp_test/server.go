package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

const (
	Cd    = "cd"
	Ls    = "ls"
	Get   = "get"
	Send  = "send"
	Close = "close"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// 起一个写携程处理连接
		go handleConnect(conn)
	}
}

func handleConnect(conn net.Conn) {
	defer conn.Close()
	nowDir := "C:/"
	// 收到连接后先把当前目录的所有文件及文件夹打印出来
	if _, err := conn.Write([]byte(getNowDir(&nowDir))); err != nil {
		log.Fatal(err)
	}
	for {
		// 读取客户端发送的命令
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading:", err.Error())
				return
			}
			break
		}
		cmd := strings.TrimSpace(string(buf[:n]))
		fmt.Println("Command received:", cmd)
		resp := handleFTPCMD(cmd, &nowDir)
		conn.Write([]byte(resp))
	}
}

func handleFTPCMD(clientReq string, nowDir *string) (response string) {
	spaceIndex := strings.Index(clientReq, " ")
	if spaceIndex == -1 {
		response = fmt.Sprintf("input [%s] is invalid!\n", clientReq)
		return response
	}
	cmd := clientReq[0:]
	dir := clientReq[strings.Index(clientReq, " ")+1:]
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		response = fmt.Sprintf("directory[%s] is not exist!\n%s", dir, getNowDir(nowDir))
		return response
	}
	switch cmd {
	case Cd:
		response = handleCd(nowDir, dir)
	case Ls:
		response = handleLs(dir)
	case Get:
	case Send:
	case Close:
	}
	return response
}

func getNowDir(nowDir *string) string {
	str := "Now dir is " + *nowDir + "\n"
	files, _ := ioutil.ReadDir(*nowDir)
	dir := "Dir: "
	file := "File: "
	for _, f := range files {
		if f.IsDir() {
			dir += f.Name() + ", "
		} else {
			file += f.Name() + ", "
		}
	}
	str += dir + "\n" + file + "\n"
	return str
}

func handleCd(nowDir *string, toDir string) string {
	if _, err := os.ReadDir(toDir); err != nil {
		return "cd failed!\n" + getNowDir(nowDir)
	}
	*nowDir = toDir
	return "cd success!\n" + getNowDir(nowDir)
}

func handleLs(dir string) string {
	if _, err := os.ReadDir(dir); err != nil {
		return "ls failed!\n" + getNowDir(&dir)
	}
	return "ls success!\n" + getNowDir(&dir)
}
