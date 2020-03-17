package http

import (
	"sync"
	"testing"
)

func TestCorps_Do(t *testing.T) {
	corps := New()

	// 同步请求
	result, err := corps.Do(Post, "http://4.shiyuesoft.com/api/uaa/oauth/login").
		AddFormParam("username", "cszy006").
		AddFormParam("password", "12345678").
		AddFormParam("phoneNumber", "").
		AddFormParam("code", "").
		SynchronousExec()
	if err != nil {
		panic(err)
	}
	t.Log(string(result.data))


	// 异步请求
	var wait sync.WaitGroup
	wait.Add(1)
	corps.Do(Post, "http://4.shiyuesoft.com/api/uaa/oauth/login").
		AddFormParam("username", "cszy006").
		AddFormParam("password", "12345678").
		AddFormParam("phoneNumber", "").
		AddFormParam("code", "").
		AsynchronousExec(func(result *Result, err error) {
			if err != nil {
				panic(err)
			}
			t.Log(string(result.data))
			wait.Done()
		})
	wait.Wait()
}