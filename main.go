// expect：be sure to finish!
// author：KercyLAN
// create at：2020-3-6 0:06

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)


func main()  {
	conn, err := net.Dial("tcp", "212.64.53.236:55014")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请发送信息(退出请输入Q):")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte("send " + trimmedInput))
		if err != nil {
			return
		}
	}
}
