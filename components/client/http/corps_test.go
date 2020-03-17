package http

import (
	"sync"
	"testing"
)

func TestCorps_Do(t *testing.T) {
	var wait sync.WaitGroup
	wait.Add(1)
	corps := New()

	// 异步请求
	corps.Do(Post, "xxx").
		AddFormParam("username", "xxx").
		AddFormParam("password", "xxx").
		AddFormParam("phoneNumber", "").
		AddFormParam("code", "").
		AsynchronousExec(func(result *Result, err error) {
			if err != nil {
				panic(err)
			}
			t.Log("异步", string(result.data))
			wait.Done()
		})

	// 同步请求
	result, err := corps.Do(Get, "xxx").
		AddUrlParamBunch(`gradeValue=&courseValue=&yearValue=&termValue=&start=1&size=30&isOnline=&_=1584428773730`).
		SynchronousExec()
	if err != nil {
		panic(err)
	}
	t.Log("同步", string(result.data))


	wait.Wait()

}