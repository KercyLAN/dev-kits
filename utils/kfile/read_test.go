// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 12:16

package kfile

import (
	"github.com/KercyLAN/dev-kits/utils/kruntime"
	"testing"
)

var filePath = kruntime.PathWork() + "\\test.txt"

func TestReadBlockHook(t *testing.T) {
	err := ReadBlockHook(filePath, 3, func(data []byte) {
		t.Log(string(data))
	})
	if err != nil {
		panic(err)
	}
}

func TestReadLineHook(t *testing.T) {
	err := ReadLineHook(filePath, func(data []byte) {
		t.Logf("%v", string(data))
	})
	if err != nil {
		panic(err)
	}
}

func TestReadOnce(t *testing.T) {
	data, err := ReadOnce(filePath)
	if err != nil {
		panic(err)
	}
	t.Logf(string(data))
}