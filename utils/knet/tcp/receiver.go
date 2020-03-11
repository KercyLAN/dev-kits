// expect：be sure to finish!
// author：KercyLAN
// create at：2020-3-5 10:16

package tcp

import (
	"fmt"
	"github.com/KercyLAN/dev-kits/utils/kencoder"
	"log"
	"net"
)

// 连接处理程序描述
type ConnHandler func(conn net.Conn, err error)

// 处理TCP连接的接收器
type Receiver struct {
	address 				string						// 监听地址
	handler 				ConnHandler					// 连接处理者
	encoder 				kencoder.Encoder			// 编码器
}

// 设置连接处理器
func (slf *Receiver) SetHandler(connHandler ConnHandler) {
	slf.handler = connHandler
}

// 设置编码器
func (slf *Receiver) SetEncoder(encoder kencoder.Encoder) {
	slf.encoder = encoder
}

// 开始接收
func (slf *Receiver) Accept() error {
	tcpListener, err := net.Listen("tcp", slf.address)
	if err != nil {
		return err
	}

	for {
		slf.handler(tcpListener.Accept())
	}
}

func (slf *Receiver) defaultConnHandler (conn net.Conn, err error) {
	defer conn.Close()
	log.Println(conn.RemoteAddr().String() + " join, but not have handler accept this connection. you can use func \"Receiver.SetConnHandler()\" set handler.")
	readBuffer := make([]byte, 512)
	for {
		readLen, err := conn.Read(readBuffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		log.Println(conn.RemoteAddr().String() + " " + slf.encoder(readBuffer[:readLen]))
		conn.Write([]byte("Welcome use \"KercyLAN TCP Server\""))
	}
}

// 构建一个TCP连接接收器实例
func NewReceiver(address string) *Receiver {
	this := &Receiver{
		address: address,
		encoder: kencoder.GbkTo.UTF8,
	}
	this.handler = this.defaultConnHandler
	return this
}
