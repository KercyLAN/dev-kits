package file

import (
	"sync"
	"testing"
)

var corps *Corps

func init() {
	conf := NewConf()
	conf.MaxFiler = 100
	conf.BufferSize = 50 * 1024 * 1024
	conf.DefaultIdle = 100
	corps = New(conf)
}

func TestCorps_Do(t *testing.T) {
	var wait sync.WaitGroup
	wait.Add(1)
	corps.Do(NewReadFileHandler(`C:\Users\Administrator\Desktop\t.log`, func(data []byte, err error) {
		if err != nil {
			panic(err)
		}
	}))
	wait.Wait()
}
