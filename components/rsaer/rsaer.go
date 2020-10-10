// Package rsaer 针对RSA加密或解密的助手
package rsaer

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

// RsaEr RSA助手数据结构
type RsaEr struct {
	publicKey 			[]byte 			// 公钥
	privateKey 			[]byte			// 私钥
}

// GetPublicKey 获取公钥
func (slf *RsaEr) GetPublicKey() []byte {
	return slf.publicKey
}

// GetPrivateKey 获取私钥
func (slf *RsaEr) GetPrivateKey() []byte {
	return slf.privateKey
}

// NewRsaEr 创建一个新的RSA助手
func NewRsaEr(publicKeyAndPrivateKey ...[]byte) *RsaEr {
	// 如果未传入或传入的参数不为两个，那么自己生成长度为1024的公私钥
	// 如果刚好是2个参数，参数1：公钥，参数2：私钥
	this := new(RsaEr)
	if len(publicKeyAndPrivateKey) == 2 {
		this.publicKey = publicKeyAndPrivateKey[0]
		this.privateKey = publicKeyAndPrivateKey[1]
	}else {
		if err := this.GenRsaKey(1024); err != nil {
			panic(errors.New(fmt.Sprint("Create rsa key failed...\r\n", err.Error())))
		}
	}
	return this
}

// RsaEncrypt Rsa加密
func (slf *RsaEr) RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(slf.publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// RsaDecrypt Rsa解密
func (slf *RsaEr) RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(slf.privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}


// GenRsaKey 发布一对新的密钥，原有密钥失效
func (slf *RsaEr) GenRsaKey(keyLength int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	priBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	slf.privateKey = pem.EncodeToMemory(priBlock)

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	slf.publicKey = pem.EncodeToMemory(publicBlock)

	if err != nil {
		return err
	}
	return nil
}