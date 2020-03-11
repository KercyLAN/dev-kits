// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 23:22

package upload

import (
	"github.com/KercyLAN/dev-kits/application"
	"github.com/KercyLAN/dev-kits/components/cloud"
	"github.com/KercyLAN/dev-kits/utils/knet"
	"net/http"
	"testing"
)

func TestFileUploadReceiver_DisposePart(t *testing.T) {
	application.Initialize(cloud.NewNode(knet.IP().String()))
	var receiver, err = New("./storage"); if err != nil {
		panic(err)
	}
	receiver.SetPartUploadBufferSize(10)

	// 创建一个分块上传任务
	receiver.NewPartTask("hash", 3)
	http.HandleFunc("/upload", func(w http.ResponseWriter,r *http.Request){
		hash, err := receiver.DisposePart(r)
		if err != nil {
			t.Log(err)
		}
		fileMeta, err := receiver.DisposePartMerge(hash, "img.jpg")
		if err != nil {
			t.Log(err)
		}else {
			t.Log(fileMeta)
		}

	})

	//监听8080端口
	http.ListenAndServe(":8080",nil)
}

func TestFileUploadReceiver_PostHandler(t *testing.T) {
	application.Initialize(cloud.NewNode(knet.IP().String()))
	var receiver, err = New("./storage"); if err != nil {
		panic(err)
	}

	receiver.SetFastUploadHandler(func(hash string) *FileMeta {
		if hash == "hash1" {
			fileMeta := NewFileMeta("激活码 (1).txt", 3097)
			fileMeta.id = "060bb006-db1b-4b00-9244-3e319ebf2cb3"
			fileMeta.hash = "hash1"
			return fileMeta
		} else {
			return nil
		}
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter,r *http.Request){
		fileMetas, faileds, err := receiver.Dispose(r)
		if err != nil {
			t.Log(err)
		}
		if len(faileds) != 0 {
			for fileName, err := range faileds {
				t.Log(fileName, err)
			}
		}
		for _, fileMeta := range fileMetas {
			t.Log(fileMeta)
		}

	})
	//监听8080端口
	http.ListenAndServe(":8080",nil)
}
