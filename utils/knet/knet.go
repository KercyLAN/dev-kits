// expect：be sure to finish!
// author：KercyLAN
// create at：2019-12-10 11:40

/**
包括了网络相关的封装及描述
 */
package knet

import (
	"log"
	"net"
)

// 返回本机的出站IP
func IP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}
